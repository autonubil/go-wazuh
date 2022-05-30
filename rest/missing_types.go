package rest

type WazuhError interface {
	Code() int32
	Title() string
	Detail() string
}

/**
Vulnerability
- cve
- name
- version
- architecture
- detection_time
- severity
- cvss2_score
- cvss3_score
- external_references
- type
- status
- condition
- title
- published
- updated

*/

// Vulnerability defines model for LogSummary.
type Vulnerability struct {
	CVE                string   `json:"cve"`
	Name               *string  `json:"name,omitempty"`
	Title              *string  `json:"title,omitempty"`
	Version            *string  `json:"version,omitempty"`
	Architecture       *string  `json:"architecture,omitempty"`
	Detection_time     *string  `json:"detection_time,omitempty"`
	Type               *string  `json:"type,omitempty"`
	Status             *string  `json:"status,omitempty"`
	Condition          *string  `json:"condition,omitempty"`
	Published          *string  `json:"published,omitempty"`
	Updated            *string  `json:"updated,omitempty"`
	ExternalReferences []string `json:"external_references,omitempty"`
	CVSS2Score         *float32 `json:"cvss2_score,omitempty"`
	CVSS3Score         *float32 `json:"cvss3_score,omitempty"`
}

// AllItemsResponseAgentsSimple defines model for AllItemsResponseAgentsSimple.
type AllItemsResponseVulnerabilities struct {
	// Embedded struct due to allOf(#/components/schemas/AllItemsResponse)
	AllItemsResponse `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// Items that successfully applied the API call action
	AffectedItems []Vulnerability `json:"affected_items"`
}

type AllItemsResponseVulnerability struct {
	// Embedded struct due to allOf(#/components/schemas/ApiResponse)
	ApiResponse `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Data *struct {
		// Embedded struct due to allOf(#/components/schemas/AllItemsResponseAgentIDs)
		AllItemsResponseVulnerabilities `yaml:",inline"`
	} `json:"data,omitempty"`
}
