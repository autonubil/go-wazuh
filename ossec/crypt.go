/*
see: https://documentation.wazuh.com/4.0/development/message-format.html
*/
package ossec

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"

	"golang.org/x/crypto/blowfish"
)

func blowfishEncrypt(ppt, key []byte) []byte {
	// create the cipher
	bfCipher, err := blowfish.NewCipher(key)
	if err != nil {
		// fix this. its okay for this tester program, but ....
		panic(err)
	}

	// make ciphertext big enough to store len(ppt)
	ciphertext := make([]byte, len(ppt))

	// create the encrypter
	ecbc := cipher.NewCBCEncrypter(bfCipher, []byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10})
	// encrypt the blocks, because block cipher
	ecbc.CryptBlocks(ciphertext, ppt)
	// return ciphertext to calling function
	return ciphertext
}

func blowfishDecrypt(ppt, key []byte) []byte {
	// create the cipher
	bfCipher, err := blowfish.NewCipher(key)
	if err != nil {
		// fix this. its okay for this tester program, but ....
		panic(err)
	}
	// pad
	pad := bfCipher.BlockSize() - (len(ppt) % bfCipher.BlockSize())
	for pad > 0 && pad < bfCipher.BlockSize() {
		ppt = append(ppt, 0)
		pad--
	}

	// make ciphertext big enough to store len(ppt)
	ciphertext := make([]byte, len(ppt))

	// create the encrypter
	ecbc := cipher.NewCBCDecrypter(bfCipher, []byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10})

	ecbc.CryptBlocks(ciphertext, ppt)
	return ciphertext
}

// Use PKCS7 to fill, IOS is also 7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func aesEncrypt(ppt, key []byte) []byte {
	// create the cipher
	aesCipher, err := aes.NewCipher(key[:32])
	if err != nil {
		// fix this. its okay for this tester program, but ....
		panic(err)
	}

	// create the encrypter
	ivBytes := []byte("FEDCBA0987654321")
	ecbc := cipher.NewCBCEncrypter(aesCipher, ivBytes)
	ppt = PKCS7Padding(ppt, ecbc.BlockSize())

	// make ciphertext big enough to store len(ppt)
	// ciphertext := make([]byte, len(ppt)+ecbc.BlockSize())
	ciphertext := make([]byte, len(ppt))

	// encrypt the blocks, because block cipher
	ecbc.CryptBlocks(ciphertext, ppt)
	// ecbc.CryptBlocks(ciphertext[len(ppt):], ivBytes)

	// return ciphertext to calling function
	return ciphertext
}

func aesDecrypt(ppt, key []byte) []byte {

	// create the cipher
	aesCipher, err := aes.NewCipher(key[:32])
	if err != nil {
		// fix this. its okay for this tester program, but ....
		panic(err)
	}
	// pad
	pad := aesCipher.BlockSize() - (len(ppt) % aesCipher.BlockSize())
	for pad > 0 && pad < aesCipher.BlockSize() {
		ppt = append(ppt, 0)
		pad--
	}

	// make ciphertext big enough to store len(ppt)
	ciphertext := make([]byte, len(ppt))

	// create the encrypter
	ecbc := cipher.NewCBCDecrypter(aesCipher, []byte("FEDCBA0987654321"))

	ecbc.CryptBlocks(ciphertext, ppt)
	return ciphertext

}

func (a *Client) decryptMessage(encMsg []byte, msgSize uint32) (string, error) {
	// 1. Safety check for minimum length
	if len(encMsg) < 5 {
		return "", NewCorruptMessage("message too short")
	}

	if encMsg[0] == '!' {
		// Use bytes.Index for efficiency and to avoid string conversion here
		endAgentID := bytes.Index(encMsg[1:], []byte("!"))
		if endAgentID == -1 {
			return "", NewCorruptMessage("missing exclamation mark")
		}
		agentID := string(encMsg[1 : endAgentID+1])
		if agentID != a.AgentID {
			return "", NewCorruptMessage("AgentID not matching")
		}
		encMsg = encMsg[endAgentID+2:]
	}

	method := EncryptionMethodBlowFish
	if len(encMsg) > 4 && string(encMsg[:4]) == "#AES" {
		method = EncryptionMethodAES
		encMsg = encMsg[4:]
	}

	if len(encMsg) == 0 || encMsg[0] != ':' {
		return "", NewCorruptMessage("missing colon")
	}
	encMsg = encMsg[1:]

	var compressed []byte
	if method == EncryptionMethodBlowFish {
		compressed = blowfishDecrypt(encMsg, []byte(a.AgentHashedKey))
	} else {
		compressed = aesDecrypt(encMsg, []byte(a.AgentHashedKey))
	}

	// 2. Remove padding
	// Ensure we don't index out of bounds if compressed is empty
	for len(compressed) > 0 && compressed[0] == '!' {
		compressed = compressed[1:]
	}

	if len(compressed) == 0 {
		return "", NewCorruptMessage("empty decrypted content")
	}

	// 3. Decompress the entire message
	b := bytes.NewReader(compressed)
	r, err := zlib.NewReader(b)
	if err != nil {
		return "", fmt.Errorf("zlib reader init: %v", err)
	}
	defer r.Close()

	// Use io.ReadAll to handle messages of any size
	decryptedMsg, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("zlib decompression failed: %v", err)
	}

	return string(decryptedMsg), nil
}

func (a *Client) cryptMsg(msg string) ([]byte, uint32) {
	msgSize := uint(len(msg))
	rand1 := rand.Intn(65536)

	if a.localCount >= 9997 {
		a.localCount = 0
		a.globalCount++
	}
	a.localCount++

	tmpMsg := fmt.Sprintf("%05d%010d:%04d:%s",
		rand1, a.globalCount, a.localCount,
		msg)

	md5Sum := fmt.Sprintf("%x", md5.Sum([]byte(tmpMsg)))
	finMsg := fmt.Sprintf("%s%s", md5Sum, tmpMsg)

	// --- REPLACED LIBDEFLATE WITH STANDARD ZLIB ---
	var b bytes.Buffer
	// Wazuh uses Level 9 (Best Compression)
	w, err := zlib.NewWriterLevel(&b, zlib.BestCompression)
	if err != nil {
		return nil, 0
	}

	_, err = w.Write([]byte(finMsg))
	if err != nil {
		return nil, 0
	}
	w.Close() // Close is required to flush the zlib trailer

	compressedMsg := b.Bytes()
	cmpSize := uint(len(compressedMsg))
	// ----------------------------------------------

	/* Pad the message (needs to be div by 8) */
	bfSize := 8 - (cmpSize % 8)
	if bfSize == 0 {
		bfSize = 8
	}

	// We recreate the padded string using the compressed bytes
	padding := "!!!!!!!!!!!!!!!!"[:bfSize]
	paddedMsg := append([]byte(padding), compressedMsg...)
	cmpSize += bfSize

	/* Get average sizes */
	a.cOrigSize += msgSize
	a.cCompSize += cmpSize

	var cryptoToken string
	var encrypted []byte
	if a.EncryptionMethod == EncryptionMethodAES {
		cryptoToken = "#AES:"
		encrypted = aesEncrypt(paddedMsg, []byte(a.AgentHashedKey))
	} else {
		cryptoToken = ":"
		encrypted = blowfishEncrypt(paddedMsg, []byte(a.AgentHashedKey))
	}

	var msgEncrypted string
	if a.AgentAllowedIPs == "any" {
		msgEncrypted = fmt.Sprintf("!%s!%s%s", a.AgentID, cryptoToken, string(encrypted))
	} else {
		msgEncrypted = fmt.Sprintf("%s%s", cryptoToken, string(encrypted))
	}

	finalData := []byte(msgEncrypted)
	finalSize := uint32(len(finalData))

	return finalData, finalSize
}

type CorruptMessage struct {
	typ string
}

func NewCorruptMessage(typ string) CorruptMessage {
	return CorruptMessage{
		typ: typ,
	}
}

func (cme CorruptMessage) Error() string {
	return fmt.Sprintf("corrupt message (%s)", cme.typ)
}
