package ossec

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"

	"github.com/autonubil/go-wazuh/pkg/zlib"
)

// SecureCounter holds the OSSEC replay counter pair (global:local) that guards
// against message replay. Each counter domain has exactly one writer.
type SecureCounter struct {
	Global uint32
	Local  uint32
}

// HashKey derives the OSSEC session key from the agent triple, matching the
// derivation in NewAgent and ossec-hids keys.c. The result is the actual
// Blowfish/AES key used on the wire (equal to Client.AgentHashedKey).
func HashKey(agentID, agentName, agentKey string) string {
	filesum1 := fmt.Sprintf("%00x", md5.Sum([]byte(agentName)))
	filesum2 := fmt.Sprintf("%00x", md5.Sum([]byte(agentID)))
	finalStr := fmt.Sprintf("%s%s", filesum1, filesum2)
	filesum1 = fmt.Sprintf("%00x", md5.Sum([]byte(finalStr)))[0:15]
	filesum2 = fmt.Sprintf("%00x", md5.Sum([]byte(agentKey)))
	return fmt.Sprintf("%s%s", filesum2, filesum1)
}

// pkcs7Unpad removes PKCS7 padding, validating the pad length against untrusted
// input so a malformed final byte cannot slice out of bounds.
func pkcs7Unpad(b []byte, blockSize int) ([]byte, error) {
	n := len(b)
	if n == 0 || n%blockSize != 0 {
		return nil, NewCorruptMessage("bad pkcs7 length")
	}
	pad := int(b[n-1])
	if pad == 0 || pad > blockSize || pad > n {
		return nil, NewCorruptMessage("bad pkcs7 padding")
	}
	return b[:n-pad], nil
}

// DecodeSecureMessage decrypts one de-framed OSSEC payload using hashedKey (the
// derived session key, see HashKey), the way a manager's remoted would. It
// returns the agentID parsed from the optional "!id!" prefix ("" when absent),
// the replay counter, and the cleartext message. Every field is bounds-checked;
// malformed input yields a CorruptMessage error rather than a panic. The caller
// is responsible for validating the counter against its own per-agent state.
func DecodeSecureMessage(hashedKey []byte, payload []byte) (agentID string, ctr SecureCounter, msg string, err error) {
	if len(payload) == 0 {
		return "", ctr, "", NewCorruptMessage("empty payload")
	}
	enc := payload

	if enc[0] == '!' {
		end := strings.IndexByte(string(enc[1:]), '!')
		if end == -1 {
			return "", ctr, "", NewCorruptMessage("missing exclamation mark")
		}
		agentID = string(enc[1 : end+1])
		enc = enc[end+2:]
	}

	method := EncryptionMethodBlowFish
	if len(enc) >= 4 && string(enc[:4]) == "#AES" {
		method = EncryptionMethodAES
		enc = enc[4:]
	}
	if len(enc) == 0 || enc[0] != ':' {
		return "", ctr, "", NewCorruptMessage("missing colon")
	}
	enc = enc[1:]
	if len(enc) == 0 {
		return "", ctr, "", NewCorruptMessage("empty ciphertext")
	}

	var compressed []byte
	if method == EncryptionMethodBlowFish {
		compressed = blowfishDecrypt(enc, hashedKey)
	} else {
		compressed = aesDecrypt(enc, hashedKey)
		// aesEncrypt applies PKCS7 padding that aesDecrypt leaves in place;
		// strip it here so only the exact zlib stream reaches the reader
		// (pkg/zlib validates the trailing bytes as the adler32 checksum).
		compressed, err = pkcs7Unpad(compressed, 16)
		if err != nil {
			return "", ctr, "", err
		}
	}
	for len(compressed) > 0 && compressed[0] == '!' {
		compressed = compressed[1:]
	}
	if len(compressed) == 0 {
		return "", ctr, "", NewCorruptMessage("empty plaintext after unpad")
	}

	r, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return "", ctr, "", err
	}
	defer r.Close()
	plain, err := io.ReadAll(r)
	if err != nil && err != io.EOF {
		return "", ctr, "", err
	}

	ctr, msg, err = parseDecoded(plain)
	return agentID, ctr, msg, err
}

// parseDecoded splits the decompressed OSSEC plaintext, which has the layout
// <md5:32hex><rand:5><global:10>:<local:4>:<msg>, verifying the MD5 integrity
// sum over everything after it.
func parseDecoded(plain []byte) (SecureCounter, string, error) {
	var ctr SecureCounter
	if len(plain) < 32 {
		return ctr, "", NewCorruptMessage("plaintext shorter than checksum")
	}
	sum := string(plain[:32])
	body := plain[32:]
	if fmt.Sprintf("%x", md5.Sum(body)) != sum {
		return ctr, "", NewCorruptMessage("md5 checksum mismatch")
	}
	// body = rand(5) global(>=1) ':' local(>=1) ':' msg
	i := bytes.IndexByte(body, ':')
	if i < 6 { // at least 5 rand digits + 1 global digit
		return ctr, "", NewCorruptMessage("missing global/local separator")
	}
	rest := body[i+1:]
	j := bytes.IndexByte(rest, ':')
	if j < 1 {
		return ctr, "", NewCorruptMessage("missing local/msg separator")
	}
	global, err := strconv.ParseUint(string(body[5:i]), 10, 32)
	if err != nil {
		return ctr, "", NewCorruptMessage("bad global counter")
	}
	local, err := strconv.ParseUint(string(rest[:j]), 10, 32)
	if err != nil {
		return ctr, "", NewCorruptMessage("bad local counter")
	}
	ctr = SecureCounter{Global: uint32(global), Local: uint32(local)}
	return ctr, string(rest[j+1:]), nil
}

// EncryptFrame produces the framed, encrypted, "!id!"-prefixed bytes for one
// message, advancing ctr. It mirrors cryptMsg but is transport-free and takes
// an explicit key and caller-owned counter, so a single caller can multiplex
// many agent identities over shared sockets. The "!id!" prefix is always
// emitted (proxied agents are registered ip=any so the manager can demultiplex).
func EncryptFrame(hashedKey []byte, agentID, msg string, ctr *SecureCounter, method EncryptionMethod) ([]byte, error) {
	if ctr == nil {
		return nil, fmt.Errorf("ossec: nil counter")
	}
	if ctr.Local >= 9997 {
		ctr.Local = 0
		ctr.Global++
	}
	ctr.Local++

	rand1 := rand.Intn(65536) //nolint:gosec // OSSEC message nonce, not a security secret
	tmpMsg := fmt.Sprintf("%05d%010d:%04d:%s", rand1, ctr.Global, ctr.Local, msg)
	sum := fmt.Sprintf("%x", md5.Sum([]byte(tmpMsg)))
	finMsg := sum + tmpMsg

	compressed, err := zlib.Compress([]byte(finMsg), 9)
	if err != nil {
		return nil, err
	}
	bfSize := 8 - (uint(len(compressed)) % 8)
	if bfSize == 0 {
		bfSize = 8
	}
	padded := append([]byte("!!!!!!!!!!!!!!!!")[:bfSize], compressed...)

	var token string
	var encrypted []byte
	if method == EncryptionMethodAES {
		token = "#AES:"
		encrypted = aesEncrypt(padded, hashedKey)
	} else {
		token = ":"
		encrypted = blowfishEncrypt(padded, hashedKey)
	}

	payload := fmt.Appendf(nil, "!%s!%s%s", agentID, token, encrypted)
	var buf bytes.Buffer
	if err := WriteFrame(&buf, payload); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
