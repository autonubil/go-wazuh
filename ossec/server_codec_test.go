package ossec

import (
	"bytes"
	crand "crypto/rand"
	"encoding/hex"
	"testing"
)

func randomKey(t *testing.T) string {
	t.Helper()
	b := make([]byte, 32)
	if _, err := crand.Read(b); err != nil {
		t.Fatal(err)
	}
	return hex.EncodeToString(b)
}

func TestHashKeyMatchesClient(t *testing.T) {
	rawKey := randomKey(t)
	c, err := NewAgent("127.0.0.1", "005", "agent005", rawKey, WithBasePath(t.TempDir()))
	if err != nil {
		t.Fatal(err)
	}
	if got := HashKey("005", "agent005", rawKey); got != c.AgentHashedKey {
		t.Fatalf("HashKey=%q want %q", got, c.AgentHashedKey)
	}
}

func decodeCompatCase(t *testing.T, method EncryptionMethod) {
	t.Helper()
	rawKey := randomKey(t)
	c, err := NewAgent("127.0.0.1", "001", "agent001", rawKey,
		WithEncryptionMethod(method),
		WithAgentAllowedIPs("any"),
		WithBasePath(t.TempDir()))
	if err != nil {
		t.Fatal(err)
	}
	const want = "1:/var/log/secure:Failed password for root"
	payload, _ := c.cryptMsg(want)

	id, ctr, msg, err := DecodeSecureMessage([]byte(c.AgentHashedKey), payload)
	if err != nil {
		t.Fatal(err)
	}
	if id != "001" {
		t.Fatalf("id=%q want 001", id)
	}
	if msg != want {
		t.Fatalf("msg=%q want %q", msg, want)
	}
	if ctr.Local == 0 {
		t.Fatalf("counter not parsed: %+v", ctr)
	}
}

func TestDecodeSecureMessageAES(t *testing.T)      { decodeCompatCase(t, EncryptionMethodAES) }
func TestDecodeSecureMessageBlowfish(t *testing.T) { decodeCompatCase(t, EncryptionMethodBlowFish) }

func encryptFrameRoundTrip(t *testing.T, method EncryptionMethod) {
	t.Helper()
	hashed := []byte(HashKey("007", "agent007", randomKey(t)))
	ctr := &SecureCounter{Global: 1, Local: 5}
	framed, err := EncryptFrame(hashed, "007", "1:loc:hello", ctr, method)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := ReadFrame(bytes.NewReader(framed), 1<<20)
	if err != nil {
		t.Fatal(err)
	}
	id, gotCtr, msg, err := DecodeSecureMessage(hashed, payload)
	if err != nil {
		t.Fatal(err)
	}
	if id != "007" || msg != "1:loc:hello" {
		t.Fatalf("id=%q msg=%q", id, msg)
	}
	if gotCtr.Local != 6 || gotCtr.Global != 1 {
		t.Fatalf("decoded ctr=%+v want {1 6}", gotCtr)
	}
	if ctr.Local != 6 {
		t.Fatalf("caller counter not advanced: %+v", ctr)
	}
}

func TestEncryptFrameRoundTripAES(t *testing.T) { encryptFrameRoundTrip(t, EncryptionMethodAES) }
func TestEncryptFrameRoundTripBlowfish(t *testing.T) {
	encryptFrameRoundTrip(t, EncryptionMethodBlowFish)
}

func TestEncryptFrameCounterRollover(t *testing.T) {
	hashed := []byte(HashKey("009", "agent009", randomKey(t)))
	ctr := &SecureCounter{Global: 3, Local: 9997}
	if _, err := EncryptFrame(hashed, "009", "x", ctr, EncryptionMethodAES); err != nil {
		t.Fatal(err)
	}
	if ctr.Global != 4 || ctr.Local != 1 {
		t.Fatalf("rollover ctr=%+v want {4 1}", ctr)
	}
}

func TestDecodeSecureMessageRejectsGarbage(t *testing.T) {
	hashed := []byte(HashKey("001", "agent001", randomKey(t)))
	if _, _, _, err := DecodeSecureMessage(hashed, []byte("!001!:not-real-ciphertext")); err == nil {
		t.Fatal("expected error decoding garbage ciphertext")
	}
	if _, _, _, err := DecodeSecureMessage(hashed, nil); err == nil {
		t.Fatal("expected error on empty payload")
	}
}
