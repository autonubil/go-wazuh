package main

import "github.com/autonubil/go-wazuh/ossec"

func main() {
	a, err := ossec.NewAgent("172.25.14.46", "666", "zbook-cze", "c588215eddc469001c41f4aa5cc306ea20fc6815e5cfe3cc666c5ccb7bc505e0")
	if err != nil {
		panic(err)
	}
	defer a.Close()
	err = a.Handshake(nil)

	if err != nil {
		panic(err)
	}
}
