package ossec

import (
	"fmt"
	"testing"
)

func TestAgentClient(t *testing.T) {
	a, err := NewAgent("172.25.14.46", "666", "zbook-cze", "c588215eddc469001c41f4aa5cc306ea20fc6815e5cfe3cc666c5ccb7bc505e0", WithEncryptionMethod(EncryptionMethodAES))
	if err != nil {
		t.Error(err)
		return
	}
	defer a.Close()
	err = a.Handshake(nil)

	if err != nil {
		t.Error(err)
		return
	}
}

func TestEncoding(t *testing.T) {
	a, err := NewAgent("172.25.14.46", "666", "zbook-cze", "c588215eddc469001c41f4aa5cc306ea20fc6815e5cfe3cc666c5ccb7bc505e0", WithEncryptionMethod(EncryptionMethodAES))
	encrypted, len := a.cryptMsg("#ping")

	fmt.Printf("Encrypted: %d: %0x\n", len, encrypted)

	decrypted, err := a.decryptMessage(encrypted, len)

	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
