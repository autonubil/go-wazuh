package main

import "github.com/autonubil/go-wazuh/ossec"

func main() {
	cfg, err := ossec.GetAgentKey("/var/ossec/etc/client.keys")
	if err != nil {
		panic(err)
	}
	a, err := ossec.NewAgent("127.0.0.1", cfg.AgentID, cfg.AgentName, cfg.AgentKey)
	if err != nil {
		panic(err)
	}
	defer a.Close()
	err = a.Connect(true)
	if err != nil {
		panic(err)
	}

	err = a.PingServer()
	if err != nil {
		panic(err)
	}
}
