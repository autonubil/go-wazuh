package indexer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	opensearch "github.com/opensearch-project/opensearch-go"
)

type Indexer struct {
	client *opensearch.Client

	IndexName string
}

func New(client *opensearch.Client) (*Indexer, error) {
	return &Indexer{
		client:    client,
		IndexName: "wazuh-states-vulnerabilities-wazuh",
	}, nil
}

type agentVulnerabilitySource struct {
	Vulnerability AgentVulnerability `json:"_source,omitempty"`
}

type agentVulnerabilityHit struct {
	Hits []agentVulnerabilitySource `json:"hits,omitempty"`
}
type osResult struct {
	Hits agentVulnerabilityHit `json:"hits,omitempty"`
}

func (i *Indexer) GetAgentVulnerabilities(agentId string, ctx context.Context) ([]AgentVulnerability, error) {

	query := fmt.Sprintf(`{
	"query": {
		"term": {
			"agent.id": "%s"
		}
	}
}`, agentId)

	res, err := i.client.Search(
		i.client.Search.WithContext(ctx),
		i.client.Search.WithIndex(i.IndexName),
		i.client.Search.WithBody(strings.NewReader(query)),
		i.client.Search.WithFrom(0),
		i.client.Search.WithSize(2500),
		i.client.Search.WithRequestCache(true),
		i.client.Search.WithSort("vulnerability.id:asc"),
	)
	if err != nil {
		return nil, err
	}

	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	raw := buf.String()

	result := osResult{}
	err = json.Unmarshal(([]byte)(raw), &result)
	if err != nil {
		return nil, err
	}

	vulns := []AgentVulnerability{}
	for _, hit := range result.Hits.Hits {
		vulns = append(vulns, hit.Vulnerability)
	}
	return vulns, nil
}
