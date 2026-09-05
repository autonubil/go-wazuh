package rest

import (
	"encoding/json"
	"testing"
)

// TestAgentUsersItemFractionalPasswordLastChange guards the fix for the Wazuh
// syscollector error "cannot unmarshal number 1787316174.6758862 into Go struct
// field ...user.password_last_change of type int64": the API returns this
// timestamp with a fractional part, so the field must be float64, not int64.
func TestAgentUsersItemFractionalPasswordLastChange(t *testing.T) {
	const body = `{"password_last_change": 1787316174.6758862, "name": "root", "id": 0}`
	var item AgentUsersItem
	if err := json.Unmarshal([]byte(body), &item); err != nil {
		t.Fatalf("decoding AgentUsersItem with fractional password_last_change: %v", err)
	}
	if item.PasswordLastChange < 1787316174 || item.PasswordLastChange >= 1787316175 {
		t.Fatalf("PasswordLastChange = %v, want ~1787316174.67", item.PasswordLastChange)
	}
	// integer timestamps must still decode (float64 accepts them)
	if err := json.Unmarshal([]byte(`{"password_last_change": 1787316174}`), &item); err != nil {
		t.Fatalf("decoding integer password_last_change: %v", err)
	}
}
