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

func TestListAgents(t *testing.T) {
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

	agents, err := c.AgentController.GetAgents(&AgentControllerGetAgentsParams{})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Get Agents TotalAffectedItems %d\n", agents.AllItemsResponse.TotalAffectedItems)
	for i, agent := range agents.AffectedItems {
		fmt.Printf(" %d: %s on %s\n", i, *agent.Id, *agent.NodeName)
	}
}

func TestAgentNotFound(t *testing.T) {
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

	agentsList := AgentsList{AgentID("53087")}
	agents, err := c.AgentController.GetAgents(&AgentControllerGetAgentsParams{AgentsList: &agentsList})
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("Get Agents TotalFailedItems %d\n", agents.AllItemsResponse.TotalFailedItems)
	for i, error := range agents.FailedItems {
		fmt.Printf(" %d: %s on %s\n", i, *error.Id, *error.Error.Message)
	}
}
