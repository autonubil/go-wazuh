package ossec

import (
	"fmt"
	"testing"
	"time"
)

func TestAgentClient(t *testing.T) {
	a, err := NewAgent("wazuh.opsanio-elk-mimosa.exanio.net", "400444", "exanio-degustator-test-t0-test-444", "UrR55N38Q5JtY73244532312222224216k9KcCTpfA2zem7U81hg26w3An47j6w9", WithAgentAllowedIPs("any"), WithEncryptionMethod(EncryptionMethodAES), WithTCP(true))
	// a, err := NewAgent("10.1.42.2", "232333", "exanio-degustator-test-t0-test-033", "UrR55N38Q5JtY73244534YKcaTKfw6N46k9KcCTpfA2zem7U81hg26w3An47j6w9", WithAgentAllowedIPs("any"), WithEncryptionMethod(EncryptionMethodAES), WithTCP(true))
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
	if err != nil {
		t.Error(err)
		return
	}
	encrypted, len := a.cryptMsg("#ping")

	fmt.Printf("Encrypted: %d: %0x\n", len, encrypted)

	decrypted, err := a.decryptMessage(encrypted, len)

	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}

func TestAw(t *testing.T) {

	a, err := NewAgent("server.agent-cluster1.prod.eu001-prod.arcticwolf.net", "3538182", "e5df31e7-a005-49de-a4ca-4a76386cc49b_3fab09e7-53ad-4a9d-8005-e1e8ab4834d8", "radvEIB3KahGdwA0vw8mT9o2LSRRMiLSdPjyDgRHoF6phqkghF2S98hRU8Nud2QC", WithAgentIP("any"), WithTCP(true), WithEncryptionMethod(EncryptionMethodAES))
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

	time.Sleep(10 * time.Minute)
}
