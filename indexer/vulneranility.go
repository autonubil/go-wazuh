package indexer

import (
	"time"
)

// Helper functions for non-string primitives
func BoolPtr(b bool) *bool           { return &b }
func Int64Ptr(i int64) *int64        { return &i }
func Float64Ptr(f float64) *float64  { return &f }
func TimePtr(t time.Time) *time.Time { return &t }

type Build struct {
	Original string `json:"original,omitempty"`
}

type AgentHostBoot struct {
	ID string `json:"id,omitempty"`
}

type AgentHostCPU struct {
	Usage *float64 `json:"usage,omitempty"`
}

type DiskIO struct {
	Bytes *int64 `json:"bytes,omitempty"`
}

type AgentHostDisk struct {
	Read  *DiskIO `json:"read,omitempty"`
	Write *DiskIO `json:"write,omitempty"`
}

type OS struct {
	Family   string `json:"family,omitempty"`
	Full     string `json:"full"`
	Kernel   string `json:"kernel"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Type     string `json:"type"`
	Version  string `json:"version"`
}

type Host struct {
	OS OS `json:"os"`
}

type Geo struct {
	CityName       string `json:"city_name,omitempty"`
	ContinentCode  string `json:"continent_code,omitempty"`
	ContinentName  string `json:"continent_name,omitempty"`
	CountryIsoCode string `json:"country_iso_code,omitempty"`
	CountryName    string `json:"country_name,omitempty"`
	Location       string `json:"location,omitempty"` // geo_point
	Name           string `json:"name,omitempty"`
	PostalCode     string `json:"postal_code,omitempty"`
	RegionIsoCode  string `json:"region_iso_code,omitempty"`
	RegionName     string `json:"region_name,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
}

type NetworkTraffic struct {
	Bytes   *int64 `json:"bytes,omitempty"`
	Packets *int64 `json:"packets,omitempty"`
}

type AgentHostNetwork struct {
	Egress  *NetworkTraffic `json:"egress,omitempty"`
	Ingress *NetworkTraffic `json:"ingress,omitempty"`
}

type HostOS struct {
	Family   string `json:"family,omitempty"`
	Full     string `json:"full,omitempty"`
	Kernel   string `json:"kernel,omitempty"`
	Name     string `json:"name,omitempty"`
	Platform string `json:"platform,omitempty"`
	Type     string `json:"type,omitempty"`
	Version  string `json:"version,omitempty"`
}

type RiskScores struct {
	CalculatedLevel     string   `json:"calculated_level,omitempty"`
	CalculatedScore     *float64 `json:"calculated_score,omitempty"`
	CalculatedScoreNorm *float64 `json:"calculated_score_norm,omitempty"`
	StaticLevel         string   `json:"static_level,omitempty"`
	StaticScore         *float64 `json:"static_score,omitempty"`
	StaticScoreNorm     *float64 `json:"static_score_norm,omitempty"`
}

type AgentHost struct {
	Architecture string            `json:"architecture,omitempty"`
	Boot         *AgentHostBoot    `json:"boot,omitempty"`
	CPU          *AgentHostCPU     `json:"cpu,omitempty"`
	Disk         *AgentHostDisk    `json:"disk,omitempty"`
	Domain       string            `json:"domain,omitempty"`
	Geo          *Geo              `json:"geo,omitempty"`
	Hostname     string            `json:"hostname,omitempty"`
	ID           string            `json:"id,omitempty"`
	IP           string            `json:"ip,omitempty"`
	MAC          string            `json:"mac,omitempty"`
	Name         string            `json:"name,omitempty"`
	Network      *AgentHostNetwork `json:"network,omitempty"`
	OS           *HostOS           `json:"os,omitempty"`
	PidNSIno     string            `json:"pid_ns_ino,omitempty"`
	Risk         *RiskScores       `json:"risk,omitempty"`
	Type         string            `json:"type,omitempty"`
	Uptime       *int64            `json:"uptime,omitempty"`
}

type Agent struct {
	Build       *Build     `json:"build,omitempty"`
	EphemeralID string     `json:"ephemeral_id,omitempty"`
	Groups      []string   `json:"groups,omitempty"`
	Host        *AgentHost `json:"host,omitempty"`
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Type        string     `json:"type,omitempty"`
	Version     string     `json:"version,omitempty"`
}

type Package struct {
	Architecture string     `json:"architecture,omitempty"`
	BuildVersion string     `json:"build_version,omitempty"`
	Checksum     string     `json:"checksum,omitempty"`
	Description  string     `json:"description,omitempty"`
	InstallScope string     `json:"install_scope,omitempty"`
	Installed    *time.Time `json:"installed,omitempty"`
	License      string     `json:"license,omitempty"`
	Name         string     `json:"name,omitempty"`
	Path         string     `json:"path,omitempty"`
	Reference    string     `json:"reference,omitempty"`
	Size         *int64     `json:"size,omitempty"`
	Type         string     `json:"type,omitempty"`
	Version      string     `json:"version,omitempty"`
}

type Scanner struct {
	Condition string `json:"condition,omitempty"`
	Reference string `json:"reference,omitempty"`
	Source    string `json:"source,omitempty"`
	Vendor    string `json:"vendor,omitempty"`
}

type Score struct {
	Base          *float64 `json:"base,omitempty"`
	Environmental *float64 `json:"environmental,omitempty"`
	Temporal      *float64 `json:"temporal,omitempty"`
	Version       string   `json:"version,omitempty"`
}

type Vulnerability struct {
	Category        string     `json:"category,omitempty"`
	Classification  string     `json:"classification,omitempty"`
	Description     string     `json:"description,omitempty"`
	DetectedAt      *time.Time `json:"detected_at,omitempty"`
	Enumeration     string     `json:"enumeration,omitempty"`
	ID              string     `json:"id,omitempty"`
	PublishedAt     *time.Time `json:"published_at,omitempty"`
	Reference       string     `json:"reference,omitempty"`
	ReportID        string     `json:"report_id,omitempty"`
	Scanner         *Scanner   `json:"scanner,omitempty"`
	Score           *Score     `json:"score,omitempty"`
	Severity        string     `json:"severity,omitempty"`
	UnderEvaluation *bool      `json:"under_evaluation,omitempty"`
}

type Cluster struct {
	Name string `json:"name,omitempty"`
	Node string `json:"node,omitempty"`
}

type Schema struct {
	Version string `json:"version,omitempty"`
}

type Wazuh struct {
	Cluster *Cluster `json:"cluster,omitempty"`
	Schema  *Schema  `json:"schema,omitempty"`
}

type AgentVulnerability struct {
	Agent         *Agent         `json:"agent,omitempty"`
	Host          *Host          `json:"host,omitempty"`
	Package       *Package       `json:"package,omitempty"`
	Vulnerability *Vulnerability `json:"vulnerability,omitempty"`
	Wazuh         *Wazuh         `json:"wazuh,omitempty"`
}
