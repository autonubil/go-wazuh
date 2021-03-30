package ossec

import (
	"encoding/xml"
	"fmt"
)

// ConvertibleBoolean xml bool values (0,no,false / 1,yes,true)
type ConvertibleBoolean bool

// UnmarshalJSON convert string to boolean
func (bit *ConvertibleBoolean) UnmarshalText(data []byte) error {
	asString := string(data)
	if asString == "1" || asString == "yes" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "no" || asString == "false" {
		*bit = false
	} else {
		return fmt.Errorf("Boolean unmarshal error: invalid input %s", asString)
	}
	return nil
}

// LocalConfig see https://documentation.wazuh.com/4.0/user-manual/reference/ossec-conf/index.html
type LocalConfig struct {
	XMLName      xml.Name     `xml:"oss_agent"`
	ClientConfig ClientConfig `xml:"client"`
}

// NewClientConfig new client config with default values set
func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		Address:       "localhost",
		Port:          1514,
		Protocol:      "udp",
		MaxRetries:    5,
		RetryInterval: 10,
		NotifyTime:    10,
		TimeReconnect: 60,
		AutoRestart:   true,
		CryptoMethod:  "blowfish",
	}
}

// ClientConfig see: https://documentation.wazuh.com/4.0/user-manual/reference/ossec-conf/client.html
type ClientConfig struct {
	XMLName xml.Name `xml:"client"`

	// Address specifies the IP address or the hostname of the Wazuh manager.
	Address string `xml:"server>address,omitempty"`

	// Port sSpecifies the port to send events to on the manager. This must match the associated listening port configured on the Wazuh manager.
	Port uint16 `xml:"server>port,omitempty"`

	// Protocol specifies the protocol to use when connecting to the manager.
	Protocol string `xml:"server>protocol,omitempty"`

	// MaxRetries number of connection retries.
	MaxRetries uint16 `xml:"server>max_retries,omitempty"`

	// RetryInterval Time interval between connection attempts (seconds).
	RetryInterval uint16 `xml:"server>retry_interval,omitempty"`

	// ConfigProfile specifies the agent.conf profile(s) to be used by the agent.
	ConfigProfile string `xml:"config-profile,omitempty"`

	// NotifyTime specifies the time in seconds between agent checkins to the manager. More frequent checkins speed up dissemination of an updated agent.conf file to the agents, but may also put an undo load on the manager if there are a large number of agents.
	NotifyTime uint16 `xml:"notify_time,omitempty"`

	// TimeReconnect specifies the time in seconds before a reconnection is attempted. This should be set to a higher number than the notify_time parameter.
	TimeReconnect uint16 `xml:"time-reconnect,omitempty"`

	// LocalIP specifies which IP address will be used to communicate with the manager when the agent has multiple network interfaces.
	LocalIP string `xml:"local_ip,omitempty"`

	//  AutoRestart toggles on and off the automatic restart of agents when a new valid configuration is received from the manager.
	AutoRestart ConvertibleBoolean `xml:"auto_restart,omitempty"`

	// CryptoMethod choose the encryption of the messages that the agent sends to the manager.
	CryptoMethod string `xml:"crypto_method,omitempty"`
}

// LoadClientConfig Load the client configuration from a fole
func LoadClientConfig(filename string) (*ClientConfig, error) {
	return nil, nil
}
