package indexer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	opensearch "github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
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

// vulnerabilityPage is one round trip of a scrolled vulnerability search.
type vulnerabilityPage struct {
	ScrollID string                `json:"_scroll_id,omitempty"`
	Hits     agentVulnerabilityHit `json:"hits,omitempty"`
}

const (
	// vulnerabilityPageSize is the hits fetched per round trip. One agent can
	// carry tens of thousands of documents, one per package and CVE.
	vulnerabilityPageSize = 1000
	// scrollKeepAlive is how long the indexer keeps the scroll between pages.
	scrollKeepAlive = time.Minute
)

// GetAgentVulnerabilities returns every vulnerability document of an agent,
// scrolling through the index page by page, in no particular order.
func (i *Indexer) GetAgentVulnerabilities(agentId string, ctx context.Context) ([]AgentVulnerability, error) {
	query, err := json.Marshal(map[string]any{
		"query": map[string]any{"term": map[string]any{"agent.id": agentId}},
	})
	if err != nil {
		return nil, err
	}

	first := func() (*vulnerabilityPage, error) {
		return decodeVulnerabilityPage(i.client.Search(
			i.client.Search.WithContext(ctx),
			i.client.Search.WithIndex(i.IndexName),
			i.client.Search.WithBody(bytes.NewReader(query)),
			i.client.Search.WithSize(vulnerabilityPageSize),
			i.client.Search.WithSort("_doc"),
			i.client.Search.WithScroll(scrollKeepAlive),
		))
	}
	next := func(scrollID string) (*vulnerabilityPage, error) {
		return decodeVulnerabilityPage(i.client.Scroll(
			i.client.Scroll.WithContext(ctx),
			i.client.Scroll.WithScrollID(scrollID),
			i.client.Scroll.WithScroll(scrollKeepAlive),
		))
	}

	vulns, scrollID, err := collectVulnerabilities(first, next)
	if scrollID != "" {
		// Best effort: an abandoned scroll expires after scrollKeepAlive anyway.
		if res, cerr := i.client.ClearScroll(i.client.ClearScroll.WithContext(ctx), i.client.ClearScroll.WithScrollID(scrollID)); cerr == nil && res != nil && res.Body != nil {
			res.Body.Close()
		}
	}
	return vulns, err
}

// collectVulnerabilities reads pages until one comes back empty. It returns
// the last scroll id it saw so the caller can release the scroll, also on
// error.
func collectVulnerabilities(first func() (*vulnerabilityPage, error), next func(scrollID string) (*vulnerabilityPage, error)) ([]AgentVulnerability, string, error) {
	vulns := []AgentVulnerability{}
	scrollID := ""
	page, err := first()
	for {
		if err != nil {
			return nil, scrollID, err
		}
		if page.ScrollID != "" {
			scrollID = page.ScrollID
		}
		if len(page.Hits.Hits) == 0 {
			return vulns, scrollID, nil
		}
		for _, hit := range page.Hits.Hits {
			vulns = append(vulns, hit.Vulnerability)
		}
		if scrollID == "" {
			return nil, "", fmt.Errorf("vulnerability search returned hits without a scroll id")
		}
		page, err = next(scrollID)
	}
}

func decodeVulnerabilityPage(res *opensearchapi.Response, err error) (*vulnerabilityPage, error) {
	if err != nil {
		return nil, err
	}
	if res == nil || res.Body == nil {
		return nil, fmt.Errorf("empty response from the indexer")
	}
	defer res.Body.Close()
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("query returned code %d (%s)", res.StatusCode, raw)
	}
	page := &vulnerabilityPage{}
	if err := json.Unmarshal(raw, page); err != nil {
		return nil, err
	}
	return page, nil
}
