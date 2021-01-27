package ossec

import (
	"strconv"
	"strings"
	"sync"

	// "compress/zlib"
	"context"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// EncryptionMethod supported transport encryption
type EncryptionMethod int

const (
	// EncryptionMethodBlowFish use BlowFish for transprot encryption
	EncryptionMethodBlowFish = EncryptionMethod(0)
	// EncryptionMethodAES use AES for transprot encryption
	EncryptionMethodAES = EncryptionMethod(1)
)

// Client allowes to handshake with the server to reach a pending state (which allowes the agent to become a group member)
type Client struct {
	Server           string
	AgentID          string
	AgentName        string
	AgentKey         string
	AgentHashedKey   string
	AgentIP          string
	Port             int
	UDP              bool
	localCount       uint
	globalCount      uint
	evtCount         uint
	cOrigSize        uint
	cCompSize        uint
	EncryptionMethod EncryptionMethod
	ctx              context.Context
	conn             net.Conn
	mx               sync.Mutex
}

// AgentOption allows setting custom parameters during construction
type AgentOption func(*Client) error

// WithContext use a custom context
func WithContext(ctx context.Context) AgentOption {
	return func(c *Client) error {
		c.ctx = ctx
		return nil
	}
}

// WithUDP use UDP as Transport
func WithUDP(udp bool) AgentOption {
	return func(c *Client) error {
		c.UDP = udp
		return nil
	}
}

// WithTCP use TCP as Transport
func WithTCP(tcp bool) AgentOption {
	return func(c *Client) error {
		c.UDP = !tcp
		return nil
	}
}

// WithPort use specific port
func WithPort(port int) AgentOption {
	return func(c *Client) error {
		c.Port = port
		return nil
	}
}

// WithEncryptionMethod specify encryption method to use
func WithEncryptionMethod(encryptionMethod EncryptionMethod) AgentOption {
	return func(c *Client) error {
		c.EncryptionMethod = encryptionMethod
		return nil
	}
}

// WithAgentIP use specific Agent IP in messages
func WithAgentIP(agentIP string) AgentOption {
	return func(c *Client) error {
		c.AgentIP = agentIP
		return nil
	}
}

// NewAgent create a new Agent for the target server
func NewAgent(server string, agentID string, agentName string, agentKey string, opts ...AgentOption) (*Client, error) {

	// key... https://github.com/ossec/ossec-hids/blob/10f6ad82f5271593c61d9ac2f14ba918e559be7b/src/os_crypto/shared/keys.c#L31
	filesum1 := fmt.Sprintf("%00x", md5.Sum([]byte(agentName)))
	filesum2 := fmt.Sprintf("%00x", md5.Sum([]byte(agentID)))
	// fmt.Printf("filesum1: %s -> %s\n", agentName, filesum1)
	// fmt.Printf("filesum2: %s -> %s\n", agentID, filesum2)

	finalStr := fmt.Sprintf("%s%s", filesum1, filesum2)
	// fmt.Printf("finalStr: %s\n", finalStr)

	filesum1 = fmt.Sprintf("%00x", md5.Sum([]byte(finalStr)))[0:15]
	filesum2 = fmt.Sprintf("%00x", md5.Sum([]byte(agentKey)))
	// fmt.Printf("filesum1: %s\n", filesum1)
	// fmt.Printf("key md5: %s\n", filesum2)

	finalStr = fmt.Sprintf("%s%s", filesum2, filesum1)
	// fmt.Printf("key: %s\n", finalStr)
	a := &Client{
		AgentID:          agentID,
		AgentKey:         agentKey,
		AgentName:        agentName,
		AgentIP:          "any",
		AgentHashedKey:   finalStr,
		Server:           server,
		Port:             1514,
		UDP:              false,
		EncryptionMethod: EncryptionMethodBlowFish,
	}

	// mutate agent and add all optional params
	for _, o := range opts {
		if err := o(a); err != nil {
			return nil, err
		}
	}

	//
	if a.ctx == nil {
		a.ctx = context.Background()
	}

	var err error
	var localAddr net.Addr
	if a.UDP {
		a.conn, err = net.Dial("udp", fmt.Sprintf("%s:%d", a.Server, a.Port))
		localAddr = a.conn.LocalAddr().(*net.UDPAddr)
	} else {
		a.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", a.Server, a.Port))
		localAddr = a.conn.LocalAddr().(*net.TCPAddr)
	}
	if err != nil {
		return nil, err
	}

	if a.AgentIP == "any" || a.AgentIP == "" {
		addr := localAddr.String()
		lastSep := strings.LastIndex(addr, ":")
		a.AgentIP = addr[:lastSep-1]
	}

	return a, nil
}

const (
	CONTROL_HEADER     = "#!-"
	EXECD_HEADER       = "execd "
	FILE_UPDATE_HEADER = "up file "
	FILE_CLOSE_HEADER  = "close file "
	HC_STARTUP         = "agent startup "
	HC_ACK             = "agent ack "
	HC_SK_DB_COMPLETED = "syscheck-db-completed"
	HC_SK_RESTART      = "syscheck restart"
	HC_REQUEST         = "req "
	HC_FIM_DB_SFS      = "fim-db-start-first-scan"
	HC_FIM_DB_EFS      = "fim-db-end-first-scan"
	HC_FIM_DB_SS       = "fim-db-start-scan"
	HC_FIM_DB_ES       = "fim-db-end-scan"
	CFGA_DB_DUMP       = "sca-dump"
	HC_SK              = "syscheck "
	HC_FIM_FILE        = "fim_file "
	HC_FIM_REGISTRY    = "fim_registry "

	LOCALFILE_MQ = '1'
	SYSLOG_MQ    = '2'
	HOSTINFO_MQ  = '3'
	SECURE_MQ    = '4'
	SYSCHECK_MQ  = '8'
	ROOTCHECK_MQ = '9'

	maxBufferSize = 1024
	timeout       = time.Duration(30 * time.Second)
)

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (a *Client) Close() error {
	return a.conn.Close()
}

// SendMessage without waiting for an answerr a message and wait for an answer
func (a *Client) WriteMessage(msg string) error {
	a.mx.Lock()
	err := a.writeMessage(msg)
	a.mx.Unlock()
	return err
}

func (a *Client) writeMessage(msg string) error {

	encryptedMsg, msgSize := a.cryptMsg(msg)
	if !a.UDP {
		// prepend a header containing message size as 4-byte little-endian unsigned integer.
		encryptedMsg = append([]byte{0, 0, 0, 0}, encryptedMsg...)
		binary.LittleEndian.PutUint32(encryptedMsg, (uint32)(msgSize))
	}

	n, err := a.conn.Write(encryptedMsg)
	if err != nil {
		return err
	}

	fmt.Printf("packet-written: bytes=%d (%d:%d:%d) -> %s \n", n, len(encryptedMsg), a.globalCount, a.localCount, msg)

	deadline := time.Now().Add(timeout)
	err = a.conn.SetReadDeadline(deadline)
	if err != nil {
		return err
	}
	return nil
}

// SendMessage send a message and wait for an answer
func (a *Client) SendMessage(msg string) (string, error) {
	a.mx.Lock()
	s, err := a.sendMessage(msg)
	a.mx.Unlock()
	return s, err
}

func (a *Client) sendMessage(msg string) (string, error) {
	err := a.writeMessage(msg)
	if err != nil {
		return "", err
	}
	buffer := make([]byte, maxBufferSize)
	nRead, err := a.conn.Read(buffer)
	if err != nil {
		return "", err
	}
	var rawMsg []byte
	var msgSize uint32
	if !a.UDP {
		msgSize = binary.LittleEndian.Uint32(buffer)
		rawMsg = buffer[4:]
	} else {
		rawMsg = buffer
		msgSize = uint32(len(buffer))
	}
	msg, err = a.decryptMessage(rawMsg[:nRead], msgSize)
	if err != nil {
		return "", err
	}

	// parse result - get counters
	msg = msg[32:]
	rand1 := msg[:5]
	globalCount, err := strconv.Atoi(msg[5:15])
	if err != nil {
		return "", err
	}
	localCount, err := strconv.Atoi(msg[16:20])
	if err != nil {
		return "", err
	}
	msg = msg[21:]
	a.localCount = uint(localCount)
	a.globalCount = uint(globalCount)
	fmt.Printf("packet-received: bytes=%d (%s:%d:%d) '%s'\n", nRead, rand1, globalCount, localCount, msg)
	return msg, nil
}

// Handshake do a handshake
func (a *Client) Handshake(ctx context.Context) error {
	a.mx.Lock()
	defer a.mx.Unlock()
	msg := fmt.Sprintf("%s%s", CONTROL_HEADER, HC_STARTUP)
	_, err := a.sendMessage(msg)
	if err != nil {
		return err
	}

	// log start startup
	msg = fmt.Sprintf("ossec: Agent started: '%s->%s' from go-wazuh.", a.AgentID, a.AgentIP)
	msg = fmt.Sprintf("%c:%s:%s", LOCALFILE_MQ, "ossec", msg)
	err = a.writeMessage(msg)
	if err != nil {
		return err
	}

	msg = fmt.Sprintf("%c:%s: Jan 26 09:05:45 n005prtg-001 PRTG: ...[PRTG Network Monitor (N005PRTG-001)] device name status down (message)", SYSLOG_MQ, "ossec")
	err = a.writeMessage(msg)
	if err != nil {
		return err
	}

	msg = fmt.Sprintf("ossec: ossec: Agent disconnected: '%s-%s'.", a.AgentName, a.AgentIP)
	msg = fmt.Sprintf("%c:%s:%s", LOCALFILE_MQ, "ossec", msg)
	err = a.writeMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
