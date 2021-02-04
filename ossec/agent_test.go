package ossec

import (
	"fmt"
	"testing"
)

func TestAgentClient(t *testing.T) {
	a, err := NewAgent("127.0.0.1", "666", "zbook-cze", "c588215eddc469001c41f4aa5cc306ea20fc6815e5cfe3cc666c5ccb7bc505e0", WithAgentAllowedIPs("any"), WithEncryptionMethod(EncryptionMethodBlowFish))
	if err != nil {
		t.Error(err)
		return
	}
	defer a.Close()
	err = a.Connect(true)

	if err != nil {
		t.Error(err)
		return
	}

	err = a.PingServer()
	if err != nil {
		t.Error(err)
		return
	}

}

func TestEncoding(t *testing.T) {
	a, err := NewAgent("172.21.218.27", "003", "myagent", "2801fb64625a4ca5523395d8ab7370dbed275a227688542493c6577c3d9fdf2c", WithAgentIP("any"), WithEncryptionMethod(EncryptionMethodAES))
	encrypted, len := a.cryptMsg("#ping")

	fmt.Printf("Encrypted: %d: %0x\n", len, encrypted)

	decrypted, err := a.decryptMessage(encrypted, len)

	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
