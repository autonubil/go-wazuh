package ossec

import (
	"bytes"
	// "compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/4kills/go-zlib"
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

func aesEncrypt(ppt, key []byte) []byte {
	// create the cipher
	bfCipher, err := aes.NewCipher(key[:32])
	if err != nil {
		// fix this. its okay for this tester program, but ....
		panic(err)
	}

	// create the encrypter
	ecbc := cipher.NewCBCEncrypter(bfCipher, []byte("FEDCBA0987654321"))
	pad := len(ppt) % ecbc.BlockSize()
	for pad != 0 {
		ppt = append(ppt, 0)
		pad = len(ppt) % ecbc.BlockSize()
	}

	// make ciphertext big enough to store len(ppt)
	ciphertext := make([]byte, len(ppt))

	// encrypt the blocks, because block cipher
	ecbc.CryptBlocks(ciphertext, ppt)
	// return ciphertext to calling function
	return ciphertext
}

func (a *Client) decryptMessage(encMsg []byte, msgSize uint32) (string, error) {
	var compressed []byte
	// fmt.Printf("%d:%d -> '%s' %0x\n", len(encMsg), msgSize, string(encMsg), encMsg)

	if encMsg[0] == '!' {
		endAgentID := strings.Index(string(encMsg[1:]), "!")
		if endAgentID == -1 {
			return "", errors.New("Corrupt message")
		}
		agentID := string(encMsg[1 : endAgentID+1])
		if agentID != a.AgentID {
			return "", errors.New("AgentID not matching")
		}
		msgSize = msgSize - uint32(endAgentID)
		encMsg = encMsg[endAgentID+2:]
	}

	method := EncryptionMethodBlowFish
	if string(encMsg[:4]) == "#AES" {
		// compressed = aesEncrypt([]byte(tmpMsg), []byte(a.AgentHashedKey))
		compressed = make([]byte, 0)
		method = EncryptionMethodAES
		encMsg = encMsg[4:]
		msgSize = msgSize - 4
	}
	if encMsg[0] != ':' {
		return "", errors.New("Corrupt message")
	}
	encMsg = encMsg[1:]
	msgSize--
	if method == EncryptionMethodBlowFish {
		compressed = blowfishDecrypt([]byte(encMsg[0:msgSize]), []byte(a.AgentHashedKey))
	} else {
		panic("Not yet supported")
	}
	for compressed[0] == '!' {
		compressed = compressed[1:]
		msgSize--
	}
	// fmt.Printf("%0x %s\n", compressed, string(compressed))
	msg := make([]byte, msgSize)
	r, _ := zlib.NewReader(nil)
	_, msg, err := r.ReadBuffer(compressed, msg)
	r.Close()
	if err != nil {
		return "", err
	}
	// fmt.Printf("%d\n", x)

	return string(msg), nil
}

func (a *Client) cryptMsg(msg string) ([]byte, uint32) {

	msgSize := uint(len(msg))
	/* Random number, take only 5 chars ~= 2^16=65536*/
	rand1 := rand.Intn(65536)

	/* Increase local and global counters */
	if a.localCount >= 9997 {
		a.localCount = 0
		a.globalCount++
	}
	a.localCount++

	tmpMsg := fmt.Sprintf("%05d%010d:%04d:%s",
		rand1, a.globalCount, a.localCount,
		msg)

	// fmt.Printf("_tmpMsg: '%s'\n", tmpMsg)

	/* Generate MD5 of the unencrypted string */

	md5Sum := fmt.Sprintf("%x", md5.Sum([]byte(tmpMsg)))
	// fmt.Printf("_tmpmsg md5: %s\n", md5Sum)
	finMsg := fmt.Sprintf("%s%s", md5Sum, tmpMsg)
	// fmt.Printf("_finMsg: '%s'\n", finMsg)
	/* Compress the message
	* We assign the first 8 bytes for padding
	 */
	var b bytes.Buffer
	w, _ := zlib.NewWriterLevel(&b, 9)
	/* written, _ := */ w.Write([]byte(finMsg))
	w.Close()
	compressedMsg := b.Bytes()
	cmpSize := uint(len(compressedMsg))

	/* Pad the message (needs to be div by 8) */
	bfSize := 8 - (cmpSize % 8)
	if bfSize == 8 {
		bfSize = 0
	}

	// fmt.Printf("compressed: <%s>: %s (%d -> %d)\n", fmt.Sprintf("%00x", md5.Sum([]byte(compressedMsg))), compressedMsg, written, cmpSize)
	// fmt.Printf("%00x", []byte(compressedMsg))
	// fmt.Printf("\n")

	tmpMsg = fmt.Sprintf("%s%s", "!!!!!!!!"[:bfSize], string(compressedMsg))
	cmpSize += bfSize
	// fmt.Printf("tmpMsg:  '%s' (%d)\n", tmpMsg, len(tmpMsg))

	/* Get average sizes */
	a.cOrigSize += msgSize
	a.cCompSize += cmpSize
	a.evtCount++

	// always send ID

	// msgEncrypted := fmt.Sprintf(":%s:%s", a.AgentID, blowfishEncrypt([]byte(tmpMsg), []byte(a.AgentHashedKey)))
	var cryptoToken string
	var encrypted []byte
	if a.EncryptionMethod == EncryptionMethodAES {
		cryptoToken = "#AES:"
		encrypted = aesEncrypt([]byte(tmpMsg), []byte(a.AgentKey))
	} else {
		cryptoToken = ":"
		encrypted = blowfishEncrypt([]byte(tmpMsg), []byte(a.AgentHashedKey))
	}

	msgEncrypted := fmt.Sprintf("!%s!%s%s", a.AgentID, cryptoToken, encrypted)

	msgSize = uint(len(msgEncrypted))
	// fmt.Printf("'%s' : %d ->  <%s> '%s' : %d\n", tmpMsg, len(tmpMsg), fmt.Sprintf("%00x", md5.Sum([]byte(msgEncrypted))), msgEncrypted, len(msgEncrypted))

	if cmpSize < uint(len(msgEncrypted)) {
		cmpSize = uint(len(msgEncrypted))
	}
	return []byte(msgEncrypted), (uint32)(cmpSize)
}
