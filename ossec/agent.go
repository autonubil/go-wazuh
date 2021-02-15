package ossec

import (
	"bytes"
	"errors"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"context"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/autonubil/go-wazuh/sysinfo"
	"go.uber.org/zap"
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
	*AgentKey
	Server           string
	Port             uint16
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
	logger           *zap.Logger
	connected        bool
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

// WithZapLogger use a custom logger
func WithZapLogger(logger *zap.Logger) AgentOption {
	return func(c *Client) error {
		c.logger = logger
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
func WithPort(port uint16) AgentOption {
	return func(c *Client) error {
		c.Port = port
		return nil
	}
}

// WithEncryptionMethod specify encryption method to use
func WithEncryptionMethod(encryptionMethod EncryptionMethod) AgentOption {
	return func(c *Client) error {
		if encryptionMethod == EncryptionMethodAES {
			return errors.New("AES is currently not supported")
		}
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

// WithAgentAllowedIPs which IPs are allwed
func WithAgentAllowedIPs(allowedIPs string) AgentOption {
	return func(c *Client) error {
		c.AgentAllowedIPs = allowedIPs
		return nil
	}
}

// NewAgent create a new Agent for the target server
func NewAgent(server string, agentID string, agentName string, agentKey string, opts ...AgentOption) (*Client, error) {

	// key... https://github.com/ossec/ossec-hids/blob/10f6ad82f5271593c61d9ac2f14ba918e559be7b/src/os_crypto/shared/keys.c#L31
	filesum1 := fmt.Sprintf("%00x", md5.Sum([]byte(agentName)))
	filesum2 := fmt.Sprintf("%00x", md5.Sum([]byte(agentID)))
	finalStr := fmt.Sprintf("%s%s", filesum1, filesum2)
	filesum1 = fmt.Sprintf("%00x", md5.Sum([]byte(finalStr)))[0:15]
	filesum2 = fmt.Sprintf("%00x", md5.Sum([]byte(agentKey)))
	finalStr = fmt.Sprintf("%s%s", filesum2, filesum1)

	a := &Client{
		AgentKey: &AgentKey{
			AgentID:         agentID,
			AgentKey:        agentKey,
			AgentName:       agentName,
			AgentAllowedIPs: "",
			AgentIP:         os.Getenv("OSSEC_AGENT_IP"),
			AgentHashedKey:  finalStr,
		},
		Server:           server,
		Port:             1514,
		UDP:              true,
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

	if a.logger != nil {
		a.logger.Debug("New ossec agent", zap.Any("agentId", a.AgentID))
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

	maxBufferSize        = 1024
	ReadWaitTimeout      = time.Duration(30 * time.Second)
	ReadImmediateTimeout = time.Duration(1 * time.Second)
)

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (a *Client) Close() error {
	if a.connected {
		msg := fmt.Sprintf("ossec: ossec: Agent disconnected: '%s-%s'.", a.AgentName, a.AgentIP)
		msg = fmt.Sprintf("%c:%s:%s", LOCALFILE_MQ, "ossec", msg)
		err := a.writeMessage(msg)
		if err != nil {
			a.connected = false
			a.conn.Close()
			return err
		}
		a.connected = false
		return a.conn.Close()
	}
	return nil
}

// WriteMessage without waiting for an answerr a message and wait for an answer
func (a *Client) WriteMessage(msg string) error {
	a.mx.Lock()
	err := a.writeMessage(msg)
	a.mx.Unlock()
	return err
}

func nullTerminatedString(data [65]int8) string {
	var buf [512]byte // Enough for a DNS name.
	for i, b := range data[:] {
		buf[i] = uint8(b)
		if b == 0 {
			return string(buf[:i])
		}
	}
	return string(buf[:])
}

func firstUpper(src string) string {
	if len(src) == 0 {
		return ""
	}
	return strings.ToUpper(src[:1]) + src[1:]
}

// PingServer send a single ping to the server
func (a *Client) PingServer() error {
	a.mx.Lock()
	defer a.mx.Unlock()
	return a.pingServer()
}

func (a *Client) pingServer() error {
	labelIP := fmt.Sprintf("#\"_agent_ip\":%s", a.AgentIP)
	agentConfigMd5 := fmt.Sprintf("%00x", md5.Sum([]byte(a.AgentName)))
	var un syscall.Utsname
	err := syscall.Uname(&un)
	if err != nil {
		return err
	}

	osInfo := sysinfo.GetOSInfo()

	osRel := osInfo.Name
	if strings.HasPrefix(strings.ToLower(osRel), strings.ToLower(osInfo.Vendor+" ")) {
		osRel = osRel[len(osInfo.Vendor)+1:]
	}
	aname := fmt.Sprintf("%s |%s |%s |%s |%s [%s|%s: %s] - %s %s",
		firstUpper(nullTerminatedString(un.Sysname)),
		nullTerminatedString(un.Nodename),
		nullTerminatedString(un.Release),
		nullTerminatedString(un.Version),
		nullTerminatedString(un.Machine),
		firstUpper(runtime.GOOS),
		runtime.GOARCH,
		osRel,
		"go-wazuh", "v0.1.0")

	labels := ""
	sharedFiles := fmt.Sprintf("%s %s\n", agentConfigMd5, "merged.mg")

	// test example:
	//"Linux |debian10 |4.19.0-9-amd64 |#1 SMP Debian 4.19.118-2+deb10u1 (2020-06-07) |x86_64 \
	// [Debian GNU/Linux|debian: 10 (buster)] - Wazuh v3.13.0 / ab73af41699f13fdd81903b5f23d8d00\nfd756ba04d9c32c8848d4608bec41251 \
	// merged.mg\n#\"_agent_ip\":192.168.0.143\n"
	msg := fmt.Sprintf("%s%s / %s\n%s%s%s\n", CONTROL_HEADER, aname, agentConfigMd5, labels, sharedFiles, labelIP)

	response, err := a.sendMessage(msg)
	if err != nil {
		return err
	}
	doneMsg := fmt.Sprintf("%s%s", CONTROL_HEADER, FILE_CLOSE_HEADER)
	ackMsg := fmt.Sprintf("%s%s", CONTROL_HEADER, HC_ACK)
	if strings.HasPrefix(string(response), "#!-up file ") {
		fieleSpecs := strings.Split(strings.Trim(response[11:], "\n \t"), " ")
		if len(fieleSpecs) == 2 {
			if a.logger != nil {
				a.logger.Debug("Receive File", zap.String("fileName", fieleSpecs[0]))
			}

			for {
				msg, err := a.readMessage(ReadWaitTimeout)
				fmt.Println(msg)
				if err != nil {
					return err
				}
				if msg == "" {
					break
				}
				if strings.HasPrefix(msg, doneMsg) {
					return nil
				}
				if msg == ackMsg {
					break
				}
			}

			if err != nil {
				return err
			}
		}

	} else {
		if string(response) != "#!-agent ack " {
			return fmt.Errorf("Not acknowledged (%s)", response)
		}
	}
	return nil
}

func (a *Client) writeMessage(msg string) error {
	encryptedMsg, msgSize := a.cryptMsg(msg)
	if !a.UDP {
		// prepend a header containing message size as 4-byte little-endian unsigned integer.
		encryptedMsg = append([]byte{0, 0, 0, 0}, encryptedMsg...)
		binary.LittleEndian.PutUint32(encryptedMsg, (uint32)(msgSize))
	}

	_, err := a.conn.Write(encryptedMsg)
	if err != nil {
		if a.logger != nil {
			a.logger.Debug("writeMessage", zap.Any("agentId", a.AgentID), zap.String("msg", msg), zap.Error(err))
		}
		return err
	}
	if a.logger != nil {
		a.logger.Debug("writeMessage", zap.Any("agentId", a.AgentID), zap.String("msg", msg))
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

// ReadMessage read next message
func (a *Client) ReadMessage(timeout time.Duration) (string, error) {
	a.mx.Lock()
	s, err := a.readMessage(timeout)
	a.mx.Unlock()
	return s, err

}

func (a *Client) sendMessage(msg string) (string, error) {
	err := a.writeMessage(msg)
	if err != nil {
		return "", err
	}
	result, err := a.readMessage(ReadWaitTimeout)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (a *Client) readMessage(timeout time.Duration) (string, error) {
	deadline := time.Now().Add(timeout)
	err := a.conn.SetReadDeadline(deadline)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	buffer := make([]byte, maxBufferSize)
	totallyRead := 0
	for true {
		nRead, err := a.conn.Read(buffer)
		if err != nil {
			return "", err
		}
		if nRead == 0 {
			break
		}
		buf.Write(buffer[:nRead])
		totallyRead += nRead
		if nRead < maxBufferSize {
			break
		}
	}

	if totallyRead == 0 {
		return "", nil
	}
	rawMsg := buf.Bytes()
	var msgSize uint32
	if !a.UDP {
		if len(rawMsg) < 4 {
			return "", errors.New("Message too short")
		}
		msgSize = binary.LittleEndian.Uint32(rawMsg)
		rawMsg = rawMsg[4:]
		totallyRead -= 4
	} else {
		msgSize = uint32(len(rawMsg))
	}
	msg, err := a.decryptMessage(rawMsg[:totallyRead], msgSize)
	if err != nil {
		return "", err
	}

	// parse result - get counters
	msg = msg[32:]
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
	// rand1 := msg[:5]
	// fmt.Printf("packet-received: bytes=%d (%s:%d:%d) '%s'\n", nRead, rand1, globalCount, localCount, msg)
	if a.logger != nil {
		a.logger.Debug("receivedMessage", zap.Any("agentId", a.AgentID), zap.String("msg", msg))
	}
	return msg, nil
}

func (a *Client) readRaw(timeout time.Duration) ([]byte, error) {
	deadline := time.Now().Add(timeout)
	err := a.conn.SetReadDeadline(deadline)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	buffer := make([]byte, maxBufferSize)
	for true {
		nRead, err := a.conn.Read(buffer)
		if nRead == 0 {
			break
		}
		if err != nil {
			return nil, err
		}
		buf.Write(buffer[:nRead])
		if nRead < maxBufferSize {
			break
		}
	}

	return buf.Bytes(), nil
}

// Connect connect and do a handshake
func (a *Client) Connect(isStartup bool) error {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.connected = false
	var err error
	var localAddr net.Addr
	if a.UDP {
		a.conn, err = net.Dial("udp", fmt.Sprintf("%s:%d", a.Server, a.Port))
		if err != nil {
			return err
		}
		localAddr = a.conn.LocalAddr().(*net.UDPAddr)
	} else {
		a.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", a.Server, a.Port))
		if err != nil {
			return err
		}
		localAddr = a.conn.LocalAddr().(*net.TCPAddr)
	}

	// determine agentIP from actual address used
	if a.AgentIP == "" {
		addr := localAddr.String()
		lastSep := strings.LastIndex(addr, ":")
		if lastSep > 0 {
			a.AgentIP = addr[:lastSep]
		} else {
			a.AgentIP = addr
		}
	}

	msg := fmt.Sprintf("%s%s", CONTROL_HEADER, HC_STARTUP)
	response, err := a.sendMessage(msg)
	if err != nil {
		return err
	}
	if string(response) != "#!-agent ack " {
		return errors.New("Not acknowledged")
	}

	if isStartup {
		// log start startup
		msg = fmt.Sprintf("ossec: Agent started: '%s->%s' from go-wazuh.", a.AgentID, a.AgentIP)
		msg = fmt.Sprintf("%c:%s:%s", LOCALFILE_MQ, "ossec", msg)
		err = a.writeMessage(msg)
		if err != nil {
			return err
		}
		err = a.pingServer()
		if err != nil {
			return err
		}
	}
	a.connected = true
	return nil
}

// AgentLoop Process messages and keep track of connection status
func (a *Client) AgentLoop() error {
	for true {
		msg, err := a.ReadMessage(ReadImmediateTimeout)
		if err == nil {
			fmt.Println(msg)
		}
		err = a.PingServer()
		if err != nil {
			a.logger.Debug("try reconnected", zap.Any("agentId", a.AgentID))
			err = a.Connect(false)
		}
		if err != nil {
			return err
		}
		time.Sleep(time.Minute * 1)
	}
	return nil
}
