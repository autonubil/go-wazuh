package sysinfo

// OS information.
type OS struct {
	Name         string `json:"name,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
	Version      string `json:"version,omitempty"`
	Release      string `json:"release,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Build        string `json:"build,omitempty"`
	Codename     string `json:"codename,omitempty"`
}
