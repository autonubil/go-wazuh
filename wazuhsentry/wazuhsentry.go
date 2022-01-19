package wazuhsentry

import (
	"encoding/gob"
	"fmt"
	"reflect"
	"time"

	"github.com/autonubil/go-wazuh/ossec"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap/zapcore"
)

type AgentTransport struct {
	channel chan *ossec.QueuePosting
	Project string
	Version string
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
}

func (t *AgentTransport) Flush(timeout time.Duration) bool {
	return true
}

var maxErrorDepth = 5

func (t *AgentTransport) Configure(options sentry.ClientOptions) {

}

var BeforeSend = func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	if hint != nil {
		if data, ok := hint.Data.(map[string]interface{}); ok {
			event.Extra, ok = data["fields"].(map[string]interface{})
			event.Level, ok = data["level"].(sentry.Level)
		}
		if hint.OriginalException != nil {
			err := hint.OriginalException
			for i := 0; i < maxErrorDepth && err != nil; i++ {
				event.Exception = append(event.Exception, sentry.Exception{
					Value:      err.Error(),
					Type:       reflect.TypeOf(err).String(),
					Stacktrace: sentry.ExtractStacktrace(err),
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
	return event
}

func (t *AgentTransport) SendEvent(event *sentry.Event) {
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

	t := &AgentTransport{
		channel: channel,
	}
	return t, nil
}

func Init(channel chan *ossec.QueuePosting, options sentry.ClientOptions) error {
	var err error
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
		data["level"] = sentry.LevelDebug
	case zapcore.InfoLevel:
		data["level"] = sentry.LevelInfo
	case zapcore.WarnLevel:
		data["level"] = sentry.LevelWarning
	case zapcore.ErrorLevel:
	case zapcore.PanicLevel:
	case zapcore.DPanicLevel:
		data["level"] = sentry.LevelError
	case zapcore.FatalLevel:
		data["level"] = sentry.LevelFatal
	default:
		data["level"] = "debug"
	}
	data["fields"] = flds
	hint := &sentry.EventHint{
		Data: data,
	}
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
