package ossec

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

// EnrollmentConfig enrolment configuration
type EnrollmentConfig struct {
	// Manager's direction or ip address
	ManagerName string
	// Port Manager's port
	Port int
	// (optional) Name of the agent. In case of NULL enrollment message will send local hostname
	AgentName string
	AgentID   string
	AgentIP   string
	// IP address or CIDR of the agent. In case of null the manager will use the source ip
	SenderIP string
	// Forces manager to use source ip
	UseSrcIP bool
	// password verification
	AuthPass string
	// Agent Key (null if not used)
	AgentKey string
	// Agent Certificate
	AgentCert stringMap
	// CA Certificate to verify server (null if not used)
	CACert string

	Groups []string

	logger *zap.Logger
}

func (c *EnrollmentConfig) SetLogger(logger *zap.Logger) {
	c.logger = logger
}

func DefaultAgentName() (string, error) {
	hostname := os.Getenv("WAZUH_AGENT_NAME")
	if hostname == "" {
		var err error
		hostname, err = os.Hostname()
		if err != nil {
			return "", err
		}
	}
	return hostname, nil
}

func InitAgent(cfg *EnrollmentConfig) (*AgentKey, error) {
	var keyFile string
	if LocalInitInfo != nil {
		keyFile = fmt.Sprintf("%s/etc/client.keys", LocalInitInfo.Directory)
	} else {
		keyFile = "/var/ossec/etc/client.keys"
	}
	//ensure path...
	path := filepath.Dir(keyFile)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			if cfg.logger != nil {
				cfg.logger.Error("register agent - create path", zap.String("authdServer", cfg.ManagerName), zap.String("agentID", cfg.AgentName), zap.String("agentIP", cfg.AgentIP), zap.String("keyfile", keyFile), zap.Error(err))
			}
			return nil, err
		}
	}

	agentKey, err := GetAgentKey(keyFile)
	keyMapValid := err == nil && agentKey != nil

	if !keyMapValid && cfg.AuthPass != "" {
		var err2 error
		// Try to register agent
		agentKey, err2 = RegisterAgent(cfg)
		curKey, err3 := GetAgentKey(keyFile)

		if err2 == nil && agentKey != nil && (err3 != nil || curKey != agentKey) {
			// persist new key
			err2 = agentKey.WriteAgentKey(keyFile)
			if err2 != nil {
				if cfg.logger != nil {
					cfg.logger.Error("register agent - write key file", zap.String("authdServer", cfg.ManagerName), zap.String("agentID", cfg.AgentName), zap.String("agentIP", cfg.AgentIP), zap.String("path", path), zap.String("keyfile", keyFile), zap.Error(err2))
				}
				return nil, err
			}
			agentKey, err = GetAgentKey(keyFile)
			if err != nil {
				return nil, err
			}
		}
	}

	if agentKey == nil {
		return nil, errors.New("no agent key configured")
	}
	return agentKey, nil
}

// NewEnrollmentConfig initialize new enrolment config
func NewEnrollmentConfig() (*EnrollmentConfig, error) {
	cfg := &EnrollmentConfig{
		ManagerName: os.Getenv("WAZUH_AUTHD_SERVER"),
		Port:        1515,
		AgentIP:     "src",
		AuthPass:    os.Getenv("WAZUH_AUTHD_PASSWORD"),
		Groups:      make([]string, 0),
	}
	var err error
	if cfg.AgentName == "" {
		hostname, err := DefaultAgentName()
		if err != nil {
			return nil, err
		}
		cfg.AgentName = hostname
	}

	if cfg.logger == nil {
		cfg.logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

// RegisterAgent register an agent via the AuthD Service
func RegisterAgent(cfg *EnrollmentConfig) (*AgentKey, error) {
	roots := x509.NewCertPool()
	insecure := true
	if cfg.CACert != "" {
		ok := roots.AppendCertsFromPEM([]byte(cfg.CACert))
		if !ok {
			return nil, errors.New("failed to add manager´s CA certificates")
		}
		insecure = false
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", cfg.ManagerName, cfg.Port), &tls.Config{
		RootCAs:            roots,
		InsecureSkipVerify: insecure,
	})
	if err != nil {
		return nil, err
	}
	var b strings.Builder
	if cfg.AuthPass != "" {
		b.WriteString(fmt.Sprintf("OSSEC PASS: %s OSSEC A:'%s'", cfg.AuthPass, cfg.AgentName))
	} else {
		b.WriteString(fmt.Sprintf("OSSEC A:'%s'", cfg.AgentName))
	}

	if len(cfg.Groups) > 0 {
		b.WriteString(fmt.Sprintf(" G:'%s'", strings.Join(cfg.Groups, ",")))
	}

	if cfg.AgentIP != "" {
		b.WriteString(fmt.Sprintf(" IP:'%s'", cfg.AgentIP))
	} else {
		b.WriteString(fmt.Sprintf(" IP:'%s'", "src"))
	}

	b.WriteString("\n")
	cfg.logger.Debug("send message", zap.String("msg", b.String()))
	_, err = conn.Write([]byte(b.String()))
	if err != nil {
		return nil, err
	}

	readBuf := make([]byte, 1024)
	read, err := conn.Read(readBuf)

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	s := strings.Trim(string(readBuf[:read]), "\n\t ")
	if strings.HasPrefix(s, "OSSEC K:'") {
		s := s[9:]
		end := strings.LastIndex(s, "'")
		if end > 0 {
			s = s[:end]
			s = strings.ReplaceAll(s, "\n", "")
			return ParseAgentKey(s)
		}
	}
	return nil, fmt.Errorf("invalid result: %s", s)
}
