package ossec

import (
	"fmt"
	"testing"
)

func TestAuthD(t *testing.T) {
	cfg, err := NewEnrollmentConfig()
	if err != nil {
		t.Error(err)
		return
	}
	cfg.AgentName = "zbook-cze"
	cfg.AgentID = "666"
	cfg.AgentKey = "c588215eddc469001c41f4aa5cc306ea20fc6815e5cfe3cc666c5ccb7bc505e0"
	cfg.AgentIP = "any"
	key, err := RegisterAgent(cfg, []string{"default"})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Result: %v\n", key)
}
