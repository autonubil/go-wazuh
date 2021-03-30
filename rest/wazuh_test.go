package rest

import (
	"fmt"
	"testing"
)

func TestStatus(t *testing.T) {
	c, err := NewClientFromEnvironment(WithInsecure(true))
	if err != nil {
		t.Error(err)
		return
	}

	err = c.Authenticate()
	if err != nil {
		t.Error(err)
		return
	}

	status, err := c.DefaultController.DefaultInfo(&DefaultControllerDefaultInfoParams{})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Connected to %s on %s\n", *status.Title, *status.Hostname)
}

func TestAgents(t *testing.T) {
	c, err := NewClientFromEnvironment(WithInsecure(true))
	if err != nil {
		t.Error(err)
		return
	}

	err = c.Authenticate()
	if err != nil {
		t.Error(err)
		return
	}

	agents, err := c.AgentsController.GetAgents(&AgentsControllerGetAgentsParams{})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Get Agents TotalAffectedItems %d\n", agents.AllItemsResponse.TotalAffectedItems)
	for i, agent := range agents.AffectedItems {
		fmt.Printf(" %d: %s on %s\n", i, *agent.Id, *agent.NodeName)
	}
}
