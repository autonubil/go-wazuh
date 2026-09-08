package ossec

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestFrameRoundTrip(t *testing.T) {
	var buf bytes.Buffer
	payload := []byte("!001!:ciphertext-bytes")
	if err := WriteFrame(&buf, payload); err != nil {
		t.Fatal(err)
	}
	if got := binary.LittleEndian.Uint32(buf.Bytes()[:4]); got != uint32(len(payload)) {
		t.Fatalf("len header = %d, want %d", got, len(payload))
	}
	out, err := ReadFrame(&buf, 1<<20)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(out, payload) {
		t.Fatalf("round-trip mismatch: %q", out)
	}
}

func TestReadFrameRejectsOversize(t *testing.T) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, uint32(2<<20)); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadFrame(&buf, 1<<20); err == nil {
		t.Fatal("expected oversize rejection")
	}
}

func TestReadFrameZeroMaxRejected(t *testing.T) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, uint32(1)); err != nil {
		t.Fatal(err)
	}
	buf.WriteByte('x')
	if _, err := ReadFrame(&buf, 0); err == nil {
		t.Fatal("expected rejection when maxLen is 0")
	}
}
