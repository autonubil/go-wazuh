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
		AgentID:         strings.Trim(items[0], " 	"),
		AgentName:       strings.Trim(items[1], " 	"),
		AgentAllowedIPs: strings.Trim(items[2], " 	"),
		AgentKey:        strings.Trim(items[3], " 	"),
	}
	return &agent, nil
}

// GetAgentKey read from Environment and if not found there, try default file
func GetAgentKey(filename string) (*AgentKey, error) {
	agentName, err := DefaultAgentName()
	if err != nil {
		return nil, err
	}
	agentID := os.Getenv("WAZUH_AGENT_ID")
	agentIP := os.Getenv("WAZUH_AGENT_IP")
	agentKey := os.Getenv("WAZUH_AGENT_KEY")

	if agentID != "" && agentName != "" && agentKey != "" {
		if agentIP == "" {
			agentIP = "any"
		}
		key := &AgentKey{
			AgentID:         agentID,
			AgentName:       agentName,
			AgentKey:        agentKey,
			AgentAllowedIPs: agentIP,
		}
		return key, nil
	}
	return GetAgentKeyFromFile(agentName, filename)
}

func GetAgentKeyFromFile(agentName string, filename string) (*AgentKey, error) {
	keyMap, err := LoadAgentKeyMap(filename)
	if err != nil {
		return nil, err
	}
	for _, key := range keyMap {
		if key.AgentName == agentName {
			return key, nil
		}
	}
	return nil, nil
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
	if a == nil {
		return errors.New("key is null")
	}
	if a.AgentID == "" {
		return errors.New("agent id is empty")
	}
	if a.AgentName == "" {
		return errors.New("agent name is empty")
	}
	if a.AgentKey == "" {
		return errors.New("agent key is empty")
	}
	if a.AgentAllowedIPs == "" {
		return errors.New("agent allowed ips is empty")
	}
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
