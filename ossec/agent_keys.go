package ossec

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/martian/log"
)

// AgentKey a single key entry
type AgentKey struct {
	AgentID         string
	AgentName       string
	AgentKey        string
	AgentHashedKey  string
	AgentAllowedIPs string
	AgentIP         string
}

// AgentKeyMap map of agents in agents key file
type AgentKeyMap map[string]*AgentKey

// ParseAgentKey parse a single key entry line
func ParseAgentKey(line string) (*AgentKey, error) {
	items := strings.Split(line, " ")
	if len(items) < 4 {
		return nil, errors.New("invalid agent key spec")
	}
	agent := AgentKey{
		AgentID: strings.Trim(items[0], " 	"),
		AgentName: strings.Trim(items[1], " 	"),
		AgentAllowedIPs: strings.Trim(items[2], " 	"),
		AgentKey: strings.Trim(items[3], " 	"),
	}
	return &agent, nil
}

// GetAgentKeyMap read from Environment and if not found there, try default file
func GetAgentKeyMap(filename string) (AgentKeyMap, error) {
	agentID := os.Getenv("WAZUH_AGENT_ID")
	agentName := os.Getenv("WAZUH_AGENT_NAME")
	agentIP := os.Getenv("WAZUH_AGENT_IP")
	agentKey := os.Getenv("WAZUH_AGENT_KEY")

	if agentID != "" && agentName != "" && agentKey != "" {
		agentMap := AgentKeyMap{}
		if agentIP == "" {
			agentIP = "any"
		}
		key := &AgentKey{
			AgentID:         agentID,
			AgentName:       agentName,
			AgentHashedKey:  agentKey,
			AgentAllowedIPs: agentIP,
		}
		agentMap[agentID] = key
		return agentMap, nil
	}
	return LoadAgentKeyMap(filename)
}

// LoadAgentKeyMap read all agent infos from a file (/var/ossec/etc/client.keys)
func LoadAgentKeyMap(filename string) (AgentKeyMap, error) {
	if filename == "" {
		filename = "/etc/client.keys"
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	agentMap := AgentKeyMap{}
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " 	")
		if strings.HasPrefix(line, "#") {
			continue
		}
		agentKey, err := ParseAgentKey(line)
		if err == nil {
			agentMap[agentKey.AgentID] = agentKey
			log.Debugf("Read agent key for: %s", line)
		}
	}
	return agentMap, nil
}

func (a *AgentKey) WriteAgentKey(filename string) error {
	if filename == "" {
		filename = "/etc/client.keys"
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "%s %s %s %s", a.AgentID, a.AgentName, a.AgentAllowedIPs, a.AgentKey)
	return err
}
