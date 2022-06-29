package wazuhsentry

import (
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"reflect"
	"strings"
	"time"

	"github.com/autonubil/go-wazuh/ossec"
	"github.com/getsentry/sentry-go"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AgentTransport struct {
	channel          chan *ossec.QueuePosting
	Project          string
	Version          string
	wrappedTransport *sentry.HTTPTransport
}

type SentryIntegrationEvent struct {
	ossec.IntegrationEvent
	Sentry interface{} `json:"sentry"`
}
type SentryEvent struct {
	*sentry.Event
	Project string `json:"project,omitempty"`
	Version string `json:"version,omitempty"`
}

func init() {
	gob.Register(SentryIntegrationEvent{})
	gob.Register(SentryEvent{})
	gob.Register(zap.Strings("reg", []string{""}).Interface)
}

func (t *AgentTransport) Flush(timeout time.Duration) bool {
	// forward to default tansport
	t.wrappedTransport.Flush(timeout)
	return true
}

var maxErrorDepth = 5

func (t *AgentTransport) Configure(options sentry.ClientOptions) {
	// forward to default tansport
	t.wrappedTransport.Configure(options)
}

func getUserFromJWT(tokenPath string, user *sentry.User) error {
	rawToken, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return err
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(string(rawToken), claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if err != nil {
		return err
	}

	if k8s, ok := claims["kubernetes.io"]; ok {
		if pod, ok := k8s.(map[string]interface{})["pod"]; ok {
			if pidID, ok := pod.(map[string]interface{})["uid"]; ok {
				user.ID = pidID.(string)
			}
			if podName, ok := pod.(map[string]interface{})["name"]; ok {
				user.Username = podName.(string)
				return nil
			}
		}
	}
	return fmt.Errorf("name not resolved")
}

func getEnvironmentFromCert(certName string) string {
	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile(certName)
	if err != nil {
		return "kubernetes"
	}
	block, _ := pem.Decode([]byte(caCert))
	if block == nil {
		return "kubernetes"
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "kubernetes"
	}
	return strings.Split(cert.Subject.CommonName, " - ")[0]
}

var BeforeSend = func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	if hint != nil {
		if data, ok := hint.Data.(map[string]interface{}); ok {
			event.Extra, ok = data["fields"].(map[string]interface{})
			// event.Level, ok = hint.
		}
		if hint.OriginalException != nil {
			err := hint.OriginalException
			for i := 0; i < maxErrorDepth && err != nil; i++ {
				stackErr := errors.WithStack(err)
				trace := sentry.ExtractStacktrace(stackErr)
				if len(trace.Frames) > 3 {
					trace.Frames = trace.Frames[0 : len(trace.Frames)-3]
				}
				event.Exception = append(event.Exception, sentry.Exception{
					Value:      err.Error(),
					Type:       reflect.TypeOf(err).String(),
					Stacktrace: trace,
				})
				switch previous := err.(type) {
				case interface{ Unwrap() error }:
					err = previous.Unwrap()
				case interface{ Cause() error }:
					err = previous.Cause()
				default:
					err = nil
				}
			}
		}
	}

	if event.Transaction == "" {
		event.Transaction = "Zap.Logger"
	}

	return event
}

func (t *AgentTransport) SendEvent(event *sentry.Event) {
	// forward to default tansport
	t.wrappedTransport.SendEvent(event)

	meta := ossec.IntegrationMeta{
		SystemName: event.ServerName,
		Protocol:   "ossec",
	}

	sentryMsg := SentryEvent{event, "", ""}

	msg := SentryIntegrationEvent{
		ossec.IntegrationEvent{
			Integration: "webhook",
			Meta:        meta,
		},
		sentryMsg,
	}

	t.channel <- &ossec.QueuePosting{
		Location:    "internal",
		ProgramName: "sentry",
		TargetQueue: ossec.LOCALFILE_MQ,
		Raw:         msg,
	}

}

func NewAgentTransport(channel chan *ossec.QueuePosting) (*AgentTransport, error) {
	if channel == nil {
		return nil, errors.New("no channel specified")
	}
	t := &AgentTransport{
		channel:          channel,
		wrappedTransport: sentry.NewHTTPTransport(),
	}
	return t, nil
}

func Init(channel chan *ossec.QueuePosting, options sentry.ClientOptions) error {
	var err error
	if options.Environment == "" {
		if _, err := os.Stat("/run/secrets/kubernetes.io/serviceaccount/ca.crt"); err == nil {
			options.Environment = getEnvironmentFromCert("/run/secrets/kubernetes.io/serviceaccount/ca.crt")
		}
	}
	options.Transport, err = NewAgentTransport(channel)
	if err != nil {
		return err
	}

	err = sentry.Init(options)
	return err
}

type SentryCore struct {
	innerCore zapcore.Core
}

func NewWrappedCore(innerCore zapcore.Core) zapcore.Core {
	return SentryCore{innerCore}
}

func (c SentryCore) Enabled(level zapcore.Level) bool {
	result := c.innerCore.Enabled(level)
	return result
}

func (c SentryCore) With(fld []zapcore.Field) zapcore.Core { return c.innerCore.With(fld) }

func (c SentryCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if ent.Level > -1 {
		ce = ce.AddCore(ent, c)
	}
	return c.innerCore.Check(ent, ce)
}

func (c SentryCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	hub := sentry.CurrentHub()
	client, scope := hub.Client(), hub.Scope()
	if client == nil || scope == nil {
		return nil
	}
	data := make(map[string]interface{})
	flds := make(map[string]interface{})
	switch entry.Level {
	case zapcore.DebugLevel:
		scope.SetLevel(sentry.LevelDebug)
	case zapcore.InfoLevel:
		scope.SetLevel(sentry.LevelInfo)
	case zapcore.WarnLevel:
		scope.SetLevel(sentry.LevelWarning)
	case zapcore.ErrorLevel:
		scope.SetLevel(sentry.LevelError)
	case zapcore.PanicLevel:
		scope.SetLevel(sentry.LevelError)
	case zapcore.DPanicLevel:
		scope.SetLevel(sentry.LevelError)
	case zapcore.FatalLevel:
		scope.SetLevel(sentry.LevelFatal)
	default:
		scope.SetLevel(sentry.LevelDebug)
	}

	data["zapEntry"] = entry
	data["fields"] = flds
	hint := &sentry.EventHint{
		Data: data,
	}

	sentryUser := sentry.User{}
	_, err := os.Stat("/run/secrets/kubernetes.io/serviceaccount/token")
	if err == nil {
		err = getUserFromJWT("/run/secrets/kubernetes.io/serviceaccount/token", &sentryUser)
	}

	if err != nil {
		user, err := user.Current()
		if err == nil {
			sentryUser.Username = user.Username
			sentryUser.ID = user.Uid
		} else {
			sentryUser.Username = "N/A"
		}
	}

	scope.SetUser(sentryUser)
	var hn string
	hn, err = os.Hostname()
	if err != nil {
		hn = "localhost"
	}
	scope.SetTransaction(fmt.Sprintf("%s@%s", sentryUser.Username, hn))

	for _, fld := range fields {
		if fld.Interface != nil {
			if fld.Type == zapcore.ErrorType {
				hint.OriginalException = fld.Interface.(error)
				if _, ok := hint.OriginalException.(ossec.QueueError); ok {
					// skip queue errors to avoid endless loop
					return nil
				}
				hint.RecoveredException = sentry.Exception{}
			} else if fld.Type == zapcore.StringerType {
				flds[fld.Key] = fld.Interface.(fmt.Stringer).String()
			} else {
				flds[fld.Key] = fld.Interface
			}
		} else {
			switch fld.Type {
			case zapcore.StringType:
				flds[fld.Key] = fld.String
			default:
				flds[fld.Key] = fld.Integer
			}
		}
	}
	client.CaptureMessage(entry.Message, hint, scope)
	return nil
}

func (c SentryCore) Sync() error { return c.innerCore.Sync() }
