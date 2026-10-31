package indexer

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func pageOf(scrollID string, ids ...string) *vulnerabilityPage {
	p := &vulnerabilityPage{ScrollID: scrollID}
	for _, id := range ids {
		v := AgentVulnerability{}
		v.Vulnerability = &Vulnerability{ID: id}
		p.Hits.Hits = append(p.Hits.Hits, agentVulnerabilitySource{Vulnerability: v})
	}
	return p
}

func TestCollectVulnerabilities(t *testing.T) {
	boom := errors.New("boom")
	tests := []struct {
		name         string
		pages        []*vulnerabilityPage // first, then each next
		failNextAt   int                  // index into pages whose fetch fails; -1 never
		wantIDs      string
		wantScrollID string
		wantErr      bool
	}{
		{name: "no hits", pages: []*vulnerabilityPage{pageOf("s1")}, failNextAt: -1, wantIDs: "[]", wantScrollID: "s1"},
		{name: "one page", pages: []*vulnerabilityPage{pageOf("s1", "CVE-1", "CVE-2"), pageOf("s1")}, failNextAt: -1, wantIDs: "[CVE-1 CVE-2]", wantScrollID: "s1"},
		{name: "several pages, scroll id rotates", pages: []*vulnerabilityPage{pageOf("s1", "CVE-1"), pageOf("s2", "CVE-2"), pageOf("s3", "CVE-3"), pageOf("s3")}, failNextAt: -1, wantIDs: "[CVE-1 CVE-2 CVE-3]", wantScrollID: "s3"},
		{name: "error on a later page keeps the scroll id", pages: []*vulnerabilityPage{pageOf("s1", "CVE-1"), nil}, failNextAt: 1, wantScrollID: "s1", wantErr: true},
		{name: "hits without a scroll id", pages: []*vulnerabilityPage{pageOf("", "CVE-1")}, failNextAt: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := 0
			var asked []string
			fetch := func() (*vulnerabilityPage, error) {
				defer func() { n++ }()
				if n == tt.failNextAt {
					return nil, boom
				}
				if n >= len(tt.pages) {
					t.Fatalf("fetched page %d, only %d exist", n, len(tt.pages))
				}
				return tt.pages[n], nil
			}
			next := func(id string) (*vulnerabilityPage, error) {
				asked = append(asked, id)
				return fetch()
			}
			got, scrollID, err := collectVulnerabilities(fetch, next)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if scrollID != tt.wantScrollID {
				t.Fatalf("scroll id = %q, want %q", scrollID, tt.wantScrollID)
			}
			if tt.wantErr {
				return
			}
			ids := []string{}
			for _, v := range got {
				ids = append(ids, v.Vulnerability.ID)
			}
			if fmt.Sprint(ids) != tt.wantIDs {
				t.Fatalf("ids = %v, want %s", ids, tt.wantIDs)
			}
			// Each next call continues the scroll the previous page named.
			for i, id := range asked {
				if want := tt.pages[i].ScrollID; id != want {
					t.Fatalf("next call %d used scroll id %q, want %q", i, id, want)
				}
			}
		})
	}
}

func TestDecodeVulnerabilityPage(t *testing.T) {
	resp := func(code int, body string) *opensearchapi.Response {
		return &opensearchapi.Response{StatusCode: code, Body: io.NopCloser(strings.NewReader(body))}
	}
	tests := []struct {
		name     string
		res      *opensearchapi.Response
		err      error
		wantErr  bool
		wantHits int
		wantID   string
	}{
		{name: "page", res: resp(200, `{"_scroll_id":"abc","hits":{"hits":[{"_source":{"vulnerability":{"id":"CVE-2012-4542"}}}]}}`), wantHits: 1, wantID: "abc"},
		{name: "empty page", res: resp(200, `{"_scroll_id":"abc","hits":{"hits":[]}}`), wantID: "abc"},
		{name: "non-200", res: resp(404, `{"error":"no such index"}`), wantErr: true},
		{name: "bad json", res: resp(200, `{`), wantErr: true},
		{name: "transport error", err: errors.New("dial"), wantErr: true},
		{name: "no response", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			page, err := decodeVulnerabilityPage(tt.res, tt.err)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if len(page.Hits.Hits) != tt.wantHits || page.ScrollID != tt.wantID {
				t.Fatalf("page = %+v", page)
			}
		})
	}
}
