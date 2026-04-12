package gowazuh

import (
	"context"

	"github.com/autonubil/go-wazuh/ossec"
	"github.com/autonubil/go-wazuh/sysinfo"
	_ "golang.org/x/mobile/bind"
)

type OSInfo struct {
	Name         string `json:"name,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
	Version      string `json:"version,omitempty"`
	Release      string `json:"release,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Build        string `json:"build,omitempty"`
	Codename     string `json:"codename,omitempty"`
}

// AgentConfig is a plain struct that gobind will turn into a Java class
type AgentConfig struct {
	Server        string
	AgentID       string
	AgentName     string
	AgentKey      string
	ClientName    string
	ClientVersion string
	UseUDP        bool
	Port          int // Using int because uint16 isn't always smooth in Java
	AgentIP       string
	BasePath      string
	AllowedIPs    string
	OSInfo        *OSInfo
}

// GetMobile is a wrapper for the incompatible Get function
func GetMobile(info *ossec.InitInfo, key string) string {
	val, _ := info.Get(key)
	return val
}

type Client struct {
	internal   *ossec.Client
	postings   chan *ossec.QueuePosting
	outChannel chan any
}

func NewClient(cfg *AgentConfig) (*Client, error) {
	// 1. Collect options based on the struct fields
	var opts []ossec.AgentOption

	if cfg.ClientName != "" {
		opts = append(opts, ossec.WithClientName(cfg.ClientName))
	}
	if cfg.ClientVersion != "" {
		opts = append(opts, ossec.WithClientVersion(cfg.ClientVersion))
	}
	if cfg.UseUDP {
		opts = append(opts, ossec.WithUDP(true))
	} else {
		opts = append(opts, ossec.WithTCP(true))
	}
	if cfg.Port > 0 {
		opts = append(opts, ossec.WithPort(uint16(cfg.Port)))
	}
	if cfg.AgentIP != "" {
		opts = append(opts, ossec.WithAgentIP(cfg.AgentIP))
	}
	if cfg.BasePath != "" {
		opts = append(opts, ossec.WithBasePath(cfg.BasePath))
	}
	if cfg.AllowedIPs != "" {
		opts = append(opts, ossec.WithAgentAllowedIPs(cfg.AllowedIPs))
	}

	if cfg.OSInfo != nil {
		osInfo := &sysinfo.OS{
			Name:         cfg.OSInfo.Name,
			Vendor:       cfg.OSInfo.Vendor,
			Version:      cfg.OSInfo.Version,
			Release:      cfg.OSInfo.Release,
			Architecture: cfg.OSInfo.Architecture,
			Build:        cfg.OSInfo.Build,
			Codename:     cfg.OSInfo.Codename,
		}
		opts = append(opts, ossec.WithOSSettings(osInfo))
	} else {
		opts = append(opts, ossec.WithOSSettings(sysinfo.GetOSInfo()))
	}

	// 2. Call the actual NewAgent with the variadic slice
	client, err := ossec.NewAgent(cfg.Server, cfg.AgentID, cfg.AgentName, cfg.AgentKey, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{internal: client}, nil
}

func (c *Client) Connect(isStartup bool) error {
	return c.internal.Connect(isStartup)
}

func (c *Client) WaitForAgentExit() error {
	if c.outChannel != nil {
		<-c.outChannel
	}
	c.outChannel = nil
	c.postings = nil
	return nil
}

func (c *Client) Close() error {
	err := c.internal.Close()
	c.outChannel = nil
	c.postings = nil
	return err
}

func (c *Client) StartAgentLoop(closeOnError bool) error {
	ctx := context.Background()
	postings, done, err := c.internal.AgentLoop(ctx, closeOnError)
	if err != nil {
		return err
	}
	c.postings = postings
	c.outChannel = done
	return err
}
