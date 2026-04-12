package main

import (
	"fmt"
	"log"
	"os"

	"github.com/autonubil/go-wazuh/ossec"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg, err := ossec.GetAgentKey("/var/ossec/etc/client.keys")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.AgentID, cfg.AgentName, cfg.AgentKey)
	srv := os.Getenv("WAZUH_SERVER")
	if srv == "" {
		srv = "127.0.0.1"
	}
	a, err := ossec.NewAgent(srv, cfg.AgentID, cfg.AgentName, cfg.AgentKey, ossec.WithTCP(true))
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
