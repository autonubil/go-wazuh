package ossec

import (
	"bufio"
	"errors"
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
		return nil, errors.New("Invalid agent key spec")
	}
	agent := AgentKey{
		AgentID: strings.Trim(items[0], " 	"),
		AgentName: strings.Trim(items[1], " 	"),
		AgentAllowedIPs: strings.Trim(items[2], " 	"),
		AgentKey: strings.Trim(items[3], " 	"),
	}
	return &agent, nil
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
