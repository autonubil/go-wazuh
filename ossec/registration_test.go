package ossec

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseEnrollRequest(t *testing.T) {
	cases := []struct {
		raw    string
		name   string
		ip     string
		pass   string
		groups []string
	}{
		{"OSSEC A:'web01'", "web01", "", "", nil},
		{"OSSEC PASS: s3cret OSSEC A:'db2' IP:'10.0.0.5' G:'linux,prod'", "db2", "10.0.0.5", "s3cret", []string{"linux", "prod"}},
	}
	for _, tc := range cases {
		req, err := ParseEnrollRequest(tc.raw)
		if err != nil {
			t.Fatalf("ParseEnrollRequest(%q): %v", tc.raw, err)
		}
		if req.Name != tc.name || req.IP != tc.ip || req.Password != tc.pass || !reflect.DeepEqual(req.Groups, tc.groups) {
			t.Fatalf("ParseEnrollRequest(%q) = %+v", tc.raw, req)
		}
	}
	if _, err := ParseEnrollRequest("garbage"); err == nil {
		t.Fatal("expected error on malformed request")
	}
}

// TestEnrollResponseRoundTrip proves FormatEnrollResponse produces exactly what
// the existing client-side parser (ParseAgentKey, as used by RegisterAgent)
// consumes — the two halves of the contract stay compatible.
func TestEnrollResponseRoundTrip(t *testing.T) {
	k := &AgentKey{AgentID: "055", AgentName: "web01", AgentAllowedIPs: "any", AgentKey: "deadbeefkey"}
	line := FormatEnrollResponse(k)

	s := strings.Trim(line, "\n\t ")
	if !strings.HasPrefix(s, "OSSEC K:'") {
		t.Fatalf("bad prefix: %q", s)
	}
	inner := s[9:strings.LastIndex(s, "'")]
	got, err := ParseAgentKey(inner)
	if err != nil {
		t.Fatal(err)
	}
	if got.AgentID != "055" || got.AgentName != "web01" || got.AgentAllowedIPs != "any" || got.AgentKey != "deadbeefkey" {
		t.Fatalf("round-trip mismatch: %+v", got)
	}
}

func TestFormatEnrollResponseDefaultsIP(t *testing.T) {
	k := &AgentKey{AgentID: "1", AgentName: "a", AgentKey: "k"}
	if got := FormatEnrollResponse(k); !strings.Contains(got, "1 a any k") {
		t.Fatalf("expected ip default any, got %q", got)
	}
}
