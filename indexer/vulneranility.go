package indexer

type Build struct {
	Original string `json:"original"`
}

type Agent struct {
	Build       Build  `json:"build,omitempty"`
	EphemeralID string `json:"ephemeral_id,omitempty"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Version     string `json:"version"`
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

type Package struct {
	Architecture string `json:"architecture"`
	BuildVersion string `json:"build_version,omitempty"`
	Checksum     string `json:"checksum,omitempty"`
	Description  string `json:"description"`
	InstallScope string `json:"install_scope,omitempty"`
	Installed    string `json:"installed,omitempty"` // assuming installed date is a string in ISO format
	License      string `json:"license,omitempty"`
	Name         string `json:"name"`
	Path         string `json:"path,omitempty"`
	Reference    string `json:"reference,omitempty"`
	Size         int64  `json:"size"`
	Type         string `json:"type"`
	Version      string `json:"version"`
}

type Scanner struct {
	Vendor string `json:"vendor"`
}

type Score struct {
		Base          float64 `json:"base"`
	Environmental float64 `json:"environmental,omitempty"`
	Temporal      float64 `json:"temporal,omitempty"`
	Version       string  `json:"version"`
}

type Vulnerability struct {
	Category       string  `json:"category"`
	Classification string  `json:"classification"`
	Description    string  `json:"description"`
	DetectedAt     string  `json:"detected_at"` // assuming dates are in string format
	Enumeration    string  `json:"enumeration"`
	ID             string  `json:"id"`
	PublishedAt    string  `json:"published_at"` // assuming dates are in string format
	Reference      string  `json:"reference"`
	ReportID       string  `json:"report_id,omitempty"`
	Scanner        Scanner `json:"scanner"`
	Score          Score   `json:"score"`
	Severity       string  `json:"severity"`
}

type Cluster struct {
	Name string `json:"name"`
	Node string `json:"node,omitempty"`
}

type Schema struct {
	Version string `json:"version"`
}

type Wazuh struct {
	Cluster Cluster `json:"cluster"`
	Schema  Schema  `json:"schema"`
}

type AgentVulnerability struct {
	Agent         Agent         `json:"agent"`
	Host          Host          `json:"host"`
	Package       Package       `json:"package"`
	Vulnerability Vulnerability `json:"vulnerability"`
	Wazuh         Wazuh         `json:"wazuh"`
}
