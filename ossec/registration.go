package ossec

import (
	"fmt"
	"regexp"
	"strings"
)

// EnrollRequest is a parsed agent enrollment (authd) request, seen from the
// manager / registration-service side. It is the server-side counterpart to the
// client-side EnrollmentConfig used by RegisterAgent.
type EnrollRequest struct {
	Name     string   // agent name (required)
	IP       string   // requested IP or CIDR; "" lets the service decide
	Password string   // optional pre-shared enrollment password
	Groups   []string // optional group memberships
}

var (
	reEnrollName  = regexp.MustCompile(`A:'([^']*)'`)
	reEnrollIP    = regexp.MustCompile(`IP:'([^']*)'`)
	reEnrollPass  = regexp.MustCompile(`OSSEC PASS:\s*(\S+)`)
	reEnrollGroup = regexp.MustCompile(`G:'([^']*)'`)
)

// ParseEnrollRequest parses an OSSEC authd enrollment request line, e.g.
//
//	OSSEC A:'web01'
//	OSSEC PASS: secret OSSEC A:'web01' IP:'10.0.0.5' G:'linux,prod'
//
// The agent name is required; IP, password, and groups are optional.
func ParseEnrollRequest(raw string) (EnrollRequest, error) {
	var req EnrollRequest
	m := reEnrollName.FindStringSubmatch(raw)
	if m == nil || m[1] == "" {
		return req, fmt.Errorf("enrollment request missing A:'name'")
	}
	req.Name = m[1]
	if m := reEnrollIP.FindStringSubmatch(raw); m != nil {
		req.IP = m[1]
	}
	if m := reEnrollPass.FindStringSubmatch(raw); m != nil {
		req.Password = m[1]
	}
	if m := reEnrollGroup.FindStringSubmatch(raw); m != nil && m[1] != "" {
		req.Groups = strings.Split(m[1], ",")
	}
	return req, nil
}

// FormatEnrollResponse builds the authd success line the agent expects and that
// ParseAgentKey consumes: "OSSEC K:'<id> <name> <ip> <key>'". The ip field is the
// key's AgentAllowedIPs, falling back to "any".
func FormatEnrollResponse(k *AgentKey) string {
	ip := k.AgentAllowedIPs
	if ip == "" {
		ip = "any"
	}
	return fmt.Sprintf("OSSEC K:'%s %s %s %s'\n", k.AgentID, k.AgentName, ip, k.AgentKey)
}

// FormatEnrollError builds the authd error line returned to a rejected agent.
func FormatEnrollError(reason string) string {
	return fmt.Sprintf("ERROR: %s\n", reason)
}

// RegistrationService is the contract a server-side enrollment backend
// implements: given a parsed request it returns the issued agent credentials —
// by relaying to a manager's authd (see RegisterAgent), minting locally, or any
// other policy. Transport (TLS listener, connection handling) is the caller's.
type RegistrationService interface {
	Register(req EnrollRequest) (*AgentKey, error)
}
