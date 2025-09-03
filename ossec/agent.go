package ossec

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"context"
	"crypto/md5"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"go.uber.org/ratelimit"
	"go.uber.org/zap"

	"github.com/autonubil/go-wazuh/sysinfo"
	"github.com/joncrlsn/dque"
	"github.com/matishsiao/goInfo"
)

// EncryptionMethod supported transport encryption
type EncryptionMethod int

const (
	// EncryptionMethodBlowFish use BlowFish for transprot encryption
	EncryptionMethodBlowFish = EncryptionMethod(0)
	// EncryptionMethodAES use AES for transprot encryption
	EncryptionMethodAES = EncryptionMethod(1)
	// maximum number of messages that can be send ber second (500 is the hard limit on the server - be gentle  )
	SendRateLimit = 450

	// time between server pings
	NotifyTime      = 5
	SysinfoInterval = 60 // each 60th  ping -> 1/h
	WazuhVersion    = "4.3.0"
)

const (
	CONTROL_HEADER     = "#!-"
	EXECD_HEADER       = "execd "
	FILE_UPDATE_HEADER = CONTROL_HEADER + "up file "
	FILE_CLOSE_HEADER  = CONTROL_HEADER + "close file "
	HC_STARTUP         = "agent startup "
	HC_ACK             = CONTROL_HEADER + "agent ack "
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

	LOCALFILE_MQ    = '1'
	SYSLOG_MQ       = '2'
	HOSTINFO_MQ     = '3'
	SECURE_MQ       = '4'
	DBSYNC_MQ       = '5'
	SYSCHECK_MQ     = '8'
	ROOTCHECK_MQ    = '9'
	MYSQL_MQ        = 'a'
	POSTGRESQL_MQ   = 'b'
	AUTH_MQ         = 'c'
	SYSCOLLECTOR_MQ = 'd'
	CISCAT_MQ       = 'e'
	WIN_EVT_MQ      = 'f'

	RIDS_DIR        = "rids"
	REMOTE_DIR      = "remote"
	WM_SYS_LOCATION = "syscollector"

	STATS_MODULE    = 11
	FTS_MODULE      = 12
	SYSCHECK_MODULE = 13
	HOSTINFO_MODULE = 15

	ROOTCHECK_MOD    = "rootcheck"
	HOSTINFO_NEW     = "hostinfo_new"
	HOSTINFO_MOD     = "hostinfo_modified"
	FIM_MOD          = "syscheck_integrity_changed"
	FIM_NEW          = "syscheck_new_entry"
	FIM_DEL          = "syscheck_deleted"
	FIM_REG_KEY_MOD  = "syscheck_registry_key_modified"
	FIM_REG_KEY_NEW  = "syscheck_registry_key_added"
	FIM_REG_KEY_DEL  = "syscheck_registry_key_deleted"
	FIM_REG_VAL_MOD  = "syscheck_registry_value_modified"
	FIM_REG_VAL_NEW  = "syscheck_registry_value_added"
	FIM_REG_VAL_DEL  = "syscheck_registry_value_deleted"
	SYSCOLLECTOR_MOD = "syscollector"
	CISCAT_MOD       = "ciscat"
	WINEVT_MOD       = "windows_eventchannel"
	SCA_MOD          = "sca"

	/* Types of events (from decoders) */
	UNKNOWN         = 0
	SYSLOG          = 1  /* syslog message */
	IDS             = 2  /* IDS alert */
	FIREWALL        = 3  /* Firewall event */
	WEBLOG          = 7  /* Apache log */
	SQUID           = 8  /* Squid log */
	DECODER_WINDOWS = 9  /* Windows log */
	HOST_INFO       = 10 /* Host information log (from nmap or similar) */
	OSSEC_RL        = 11 /* OSSEC rule */

	maxBufferSize        = 1024 * 1024 * 2
	ReadWaitTimeout      = time.Duration(30 * time.Second)
	ReadImmediateTimeout = time.Duration(1 * time.Second)
)

// Client allowes to handshake with the server to reach a pending state (which allowes the agent to become a group member)
type Client struct {
	*AgentKey
	Server             string
	Port               uint16
	UDP                bool
	EncryptionMethod   EncryptionMethod
	ClientName         string
	ClientVersion      string
	ConfigHash         string
	RemoteFiles        map[string]RemoteFileInfo
	CurrentRemoteFile  *RemoteFileInfo
	Scanner            *SysCollector
	basePath           string
	remotePath         string
	ridsPath           string
	localCount         uint
	globalCount        uint
	evtCount           uint64
	sentCount          uint64
	receivedCount      uint64
	cOrigSize          uint
	cCompSize          uint
	sessionID          int64
	sentBytes          uint64
	sentBytesTotal     uint64
	receivedBytes      uint64
	receivedBytesTotal uint64
	ctx                context.Context
	conn               net.Conn
	mx                 sync.Mutex
	logger             *zap.Logger
	connected          bool
	rateLimit          ratelimit.Limiter
	outChannel         chan any
	un                 *goInfo.GoInfoObject
	osInfo             *sysinfo.OS
}

type FileUpdatedEvent struct {
	FileInfo *RemoteFileInfo
}

type AgentShutDownEvent struct {
}
type RemoteFileInfo struct {
	Filename string
	Hash     string
	Content  *bytes.Buffer
}

func init() {
	gob.Register(map[string]any{})
	gob.Register(QueuePosting{})
	gob.Register(FileUpdatedEvent{})
	gob.Register(AgentShutDownEvent{})
	gob.Register(RemoteFileInfo{})
	gob.Register(Client{})
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

// WithClientName use a custom client name
func WithClientName(clientName string) AgentOption {
	return func(c *Client) error {
		c.ClientName = clientName
		return nil
	}
}

// WithClientVersion use a custom client version
func WithClientVersion(clientVersion string) AgentOption {
	return func(c *Client) error {
		c.ClientVersion = clientVersion
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
		c.EncryptionMethod = encryptionMethod
		return nil
	}
}

// WithConfigHash specify a local config hash
func WithConfigHash(configHash string) AgentOption {
	return func(c *Client) error {
		c.ConfigHash = configHash
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

// WithBasePath use specific where to cache downloaded files
func WithBasePath(basePath string) AgentOption {
	return func(c *Client) error {
		c.basePath = basePath
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
	un, err := goInfo.GetInfo()
	if err != nil {
		return nil, err
	}

	a := &Client{
		AgentKey: &AgentKey{
			AgentID:         agentID,
			AgentKey:        agentKey,
			AgentName:       agentName,
			AgentAllowedIPs: "",
			AgentIP:         os.Getenv("OSSEC_AGENT_IP"),
			AgentHashedKey:  finalStr,
		},
		ConfigHash:       "",
		Server:           server,
		Port:             1514,
		UDP:              true,
		EncryptionMethod: EncryptionMethodAES,
		ClientName:       "go-wazuh",
		ClientVersion:    "v1.0.0",
		rateLimit:        ratelimit.New(SendRateLimit), // per second
		RemoteFiles:      make(map[string]RemoteFileInfo),
		un:               &un,
		osInfo:           sysinfo.GetOSInfo(),
	}

	a.Scanner = NewScanner(a)

	// mutate agent and add all optional params
	for _, o := range opts {
		if err := o(a); err != nil {
			return nil, err
		}
	}

	if a.logger == nil {
		a.logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	}

	//
	if a.ctx == nil {
		a.ctx = context.Background()
	}

	if a.logger != nil {
		a.logger.Debug("newAgent", zap.Any("agentId", a.AgentID))
	}

	if a.basePath == "" {
		if _, err := os.Stat("/var/tmp"); !os.IsNotExist(err) {
			a.basePath = "/var/tmp"
		} else if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
			a.basePath = dir
		} else {
			a.basePath = os.TempDir()
		}
	}

	a.remotePath = filepath.Join(a.basePath, REMOTE_DIR)
	if _, err := os.Stat(a.remotePath); os.IsNotExist(err) {
		os.MkdirAll(a.remotePath, 0770)
	}

	a.ridsPath = filepath.Join(a.basePath, RIDS_DIR)
	if _, err := os.Stat(a.ridsPath); os.IsNotExist(err) {
		os.MkdirAll(a.ridsPath, 0770)
	}

	err = a.ReadClientCounter()
	if err != nil {
		return nil, err
	}

	AgentCollector.Register(a)
	return a, nil
}

func (a *Client) IsConencted() bool {
	return a.connected
}

func (a *Client) GetBasePath() string {
	return a.basePath
}

// WriteClientCounter persist current counters
func (a *Client) WriteClientCounter() error {
	ridsFile := filepath.Join(a.ridsPath, a.AgentID)
	file, err := os.Create(ridsFile)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%d:%d", a.globalCount, a.localCount))
	return nil
}

// ReadClientCounter read counters from disk
func (a *Client) ReadClientCounter() error {
	ridsFile := filepath.Join(a.ridsPath, a.AgentID)
	_, err := os.Stat(ridsFile)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	file, err := os.Open(ridsFile)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	if scanner.Scan() {
		strVal := scanner.Text()
		if err := scanner.Err(); err != nil {
			return err
		}
		parts := strings.Split(strVal, ":")
		var gc int
		var lc int
		if len(parts) == 2 {
			gc, err = strconv.Atoi(parts[0])
			if err == nil {
				lc, err = strconv.Atoi(parts[1])
			}
		}

		if err != nil {
			return err
		}
		a.globalCount = uint(gc)
		a.localCount = uint(lc)
		return nil
	}
	return nil
}

func (a *Client) close(sendCloseMsg bool) error {
	if a.connected {
		if sendCloseMsg {
			a.logger.Info("managerDisconnected", zap.Any("agentId", a.AgentID))
			msg := fmt.Sprintf("ossec: ossec: Agent disconnected: '%s-%s'.", a.AgentName, a.AgentIP)
			msg = fmt.Sprintf("%c:%s:%s", LOCALFILE_MQ, "ossec", msg)
			err := a.writeMessage(msg)
			if err != nil {
				a.connected = false
				AgentCollector.Disconnect(a)
				return err
			}
		}
		a.connected = false
		a.sentBytes = 0
		a.sessionID = 0
		a.receivedBytes = 0
		AgentCollector.Disconnect(a)
		return a.conn.Close()
	}
	a.receivedBytes = 0
	a.sentBytes = 0
	a.sessionID = 0
	AgentCollector.Disconnect(a)
	return nil
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (a *Client) Close() error {
	err := a.close(true)
	AgentCollector.Unregister(a)
	return err
}

// WriteMessage without waiting for an answerr a message and wait for an answer
func (a *Client) WriteMessage(msg string) error {
	a.mx.Lock()
	err := a.writeMessage(msg)
	a.mx.Unlock()
	return err
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

func (a *Client) handleControlResponse(response string) error {
	a.logger.Debug("controlMsg", zap.Any("agentId", a.AgentID), zap.String("message", strings.Split(response, "\n")[0]))
	if strings.HasPrefix(string(response), FILE_UPDATE_HEADER) {
		fieleSpecs := strings.Split(strings.Trim(response[11:], "\n \t"), " ")
		if len(fieleSpecs) == 2 {
			if existingFile, ok := a.RemoteFiles[fieleSpecs[1]]; ok && existingFile.Hash == fieleSpecs[0] {
				return nil
			}
			if a.logger != nil {
				a.logger.Debug("receiveFile", zap.Any("agentId", a.AgentID), zap.String("fileName", fieleSpecs[1]))
			}
			a.CurrentRemoteFile = &RemoteFileInfo{
				Filename: fieleSpecs[1],
				Hash:     fieleSpecs[0],
				Content:  bytes.NewBuffer(nil),
			}
		}
		return nil
	} else if strings.HasPrefix(string(response), FILE_CLOSE_HEADER) {
		if a.logger != nil {
			a.logger.Debug("fileReceived", zap.Any("agentId", a.AgentID), zap.Int("len", a.CurrentRemoteFile.Content.Len()))
		}
		a.cacheFileHash(a.CurrentRemoteFile.Filename, a.CurrentRemoteFile.Hash, a.CurrentRemoteFile.Content.String())
		a.outChannel <- &FileUpdatedEvent{a.CurrentRemoteFile}
		a.CurrentRemoteFile = nil
		return nil
	} else if string(response) == HC_ACK {
		return nil
	} else {
		if a.CurrentRemoteFile != nil {
			// close any open file
			a.CurrentRemoteFile = nil
		}
		a.logger.Warn("unhandledControlMessage", zap.Any("agentId", a.AgentID), zap.String("msg", string(response)))
	}
	return nil
}

func (a *Client) handleResponse(response string) error {
	if strings.HasPrefix(string(response), CONTROL_HEADER) {
		return a.handleControlResponse(response)
	} else {
		if a.CurrentRemoteFile != nil {
			a.CurrentRemoteFile.Content.WriteString(response)
			a.cacheFileHash(a.CurrentRemoteFile.Filename, a.CurrentRemoteFile.Hash, a.CurrentRemoteFile.Content.String())
			return nil
		} else {
			if a.logger != nil {
				a.logger.Debug("receivedMessage", zap.Any("agentId", a.AgentID), zap.String("msg", string(response)), zap.Uint("globalCount", a.globalCount), zap.Uint("localCount", a.localCount), zap.Uint64("evtCount", a.evtCount), zap.Uint64("sentCount", a.sentCount), zap.Uint64("receivedCount", a.receivedCount))
			}
		}
	}
	return nil
}

func (a *Client) getSharedFiles() string {
	b := strings.Builder{}
	searchPath := filepath.Join(a.remotePath, "merged.mg") //"*")
	files, err := filepath.Glob(searchPath)
	if err == nil && len(files) > 0 {
		for _, filename := range files {
			content, err := os.ReadFile(filename)
			if err != nil {
				a.logger.Warn("corruptFile", zap.Any("agentId", a.AgentID), zap.String("filename", filename), zap.Error(err))
				os.Remove(filename)
				continue
			}
			h := md5.New()
			h.Write(content)
			fileInfo := RemoteFileInfo{
				Filename: filename[len(a.remotePath)+1:],
				Content:  bytes.NewBuffer([]byte(content)),
				Hash:     fmt.Sprintf("%x", h.Sum(nil)),
			}
			a.RemoteFiles[filename] = fileInfo
			b.WriteString(fmt.Sprintf("%s %s\n", fileInfo.Hash, fileInfo.Filename))
		}
	} else {
		// new agent registration. empty mergerd.mg
		b.WriteString("x merged.mg\n")
	}
	return b.String()
}

func (a *Client) cacheFileHash(fileName string, hash string, content string) error {
	fullpath := filepath.Join(a.remotePath, fileName)
	err := os.WriteFile(fullpath, []byte(content), 0644)
	if err == nil {
		a.RemoteFiles[fileName] = RemoteFileInfo{
			Filename: fileName,
			Content:  bytes.NewBuffer([]byte(content)),
			Hash:     hash,
		}
	}
	return err
}

func (a *Client) pingServer() error {
	labelIP := fmt.Sprintf("#\"_agent_ip\":%s", a.AgentIP)
	// deprecated? labelClientName := fmt.Sprintf("\"client_name\":%s\n", a.ClientName)
	// labelWazuhVersion := fmt.Sprintf("#\"_wazuh_version\":%s\n", WazuhVersion)
	// labelNodeName := fmt.Sprintf("#\"_node_name\":%s\n", a.AgentName)

	osRel := a.osInfo.Name
	if strings.HasPrefix(strings.ToLower(osRel), strings.ToLower(a.osInfo.Vendor+" ")) {
		osRel = osRel[len(a.osInfo.Vendor)+1:]
	}
	codename := ""
	if a.osInfo.Codename != "" {
		codename = fmt.Sprintf("(%s) ", a.osInfo.Codename)
	}
	aname := fmt.Sprintf("%s |%s |%s |%s |%s [%s|%s: %s%s] - %s %s",
		firstUpper(a.un.OS),
		a.un.Hostname,
		a.un.Core,
		a.un.Kernel,
		a.osInfo.Build,
		firstUpper(runtime.GOOS),
		runtime.GOARCH,
		codename,
		osRel,
		a.ClientName, a.ClientVersion)

	sharedFiles := a.getSharedFiles()

	// test example:
	//"Linux |debian10 |4.19.0-9-amd64 |#1 SMP Debian 4.19.118-2+deb10u1 (2020-06-07) |x86_64 \
	// [Debian GNU/Linux|debian: 10 (buster)] - Wazuh v3.13.0 / ab73af41699f13fdd81903b5f23d8d00\nfd756ba04d9c32c8848d4608bec41251 \
	// merged.mg\n#\"_agent_ip\":192.168.0.143\n"

	// Linux |exanio-nubo-test-m0-kula-001 |5.4.0-74-generic |#83-Ubuntu SMP Sat May 8 02:35:39 UTC 2021 |x86_64 [Ubuntu|ubuntu: 20.04.2 LTS (Focal Fossa)] - Wazuh v4.2.5 / ab73af41699f13fdd81903b5f23d8d00
	// 2c45c95db2954d2c7d0ea533f09e81a5 merged.mg
	// #"_agent_ip":10.160.220.10
	// '

	var msg string
	if a.ConfigHash != "" {
		msg = fmt.Sprintf("%s%s / %s\n%s%s\n", CONTROL_HEADER, aname, a.ConfigHash, sharedFiles, labelIP)
	} else {
		msg = fmt.Sprintf("%s%s\n%s%s\n", CONTROL_HEADER, aname, sharedFiles, labelIP)
	}
	err := a.sendMessage(msg, ReadWaitTimeout)
	if err != nil {
		return err
	}

	return nil
}

func (a *Client) writeMessage(msg string) error {
	if a.conn == nil {
		return fmt.Errorf("client is not connected")
	}

	encryptedMsg, msgSize := a.cryptMsg(msg)
	a.evtCount++

	if !a.UDP {
		// prepend a header containing message size as 4-byte little-endian unsigned integer.
		sizeBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(sizeBytes, (uint32)(msgSize))
		encryptedMsg = append(sizeBytes, encryptedMsg...)
	}

	// ensure ratelimit is honored
	prev := time.Now()
	now := a.rateLimit.Take()
	written, err := a.conn.Write(encryptedMsg)
	a.sentBytes += uint64(written)
	a.sentBytesTotal += uint64(written)
	AgentCollector.BytesSent(a, written)

	if err != nil || written == 0 {
		AgentCollector.MessageError(a, 1)
		if a.logger != nil {
			a.logger.Warn("writeMessage", zap.Any("agentId", a.AgentID), zap.String("msg", msg), zap.Int("written", written), zap.Uint64("sentBytes", a.sentBytes), zap.Uint64("sentBytesTotal", a.sentBytesTotal), zap.Duration("rateWait", now.Sub(prev)), zap.Uint("globalCount", a.globalCount), zap.Uint("localCount", a.localCount), zap.Uint64("evtCount", a.evtCount), zap.Uint64("sentCount", a.sentCount), zap.Uint64("receivedCount", a.receivedCount), zap.String("error", err.Error()))
		}
		err2 := a.close(false)
		if err2 != nil {
			a.logger.Warn("closeFailed", zap.Any("agentId", a.AgentID), zap.String("error", err.Error()))
		}
		return err
	}
	a.sentCount++
	AgentCollector.MessagesSent(a, 1)
	// time.Sleep(25 * time.Millisecond)

	if a.logger != nil {
		a.logger.Debug("writeMessage", zap.Any("agentId", a.AgentID), zap.String("msg", msg), zap.Int("written", written), zap.Uint64("sentBytes", a.sentBytes), zap.Uint64("sentBytesTotal", a.sentBytesTotal), zap.Duration("rateWait", now.Sub(prev)), zap.Uint("globalCount", a.globalCount), zap.Uint("localCount", a.localCount), zap.Uint64("evtCount", a.evtCount), zap.Uint64("sentCount", a.sentCount), zap.Uint64("receivedCount", a.receivedCount))
	}
	return nil
}

// SendMessage send a message and wait for an answer
func (a *Client) SendMessage(msg string, readTimeout time.Duration) error {
	a.mx.Lock()
	err := a.sendMessage(msg, readTimeout)
	a.mx.Unlock()
	return err
}

// ReadServerResponse read next message
func (a *Client) ReadServerResponse(timeout time.Duration) error {
	a.mx.Lock()
	err := a.readServerResponse(timeout)
	a.mx.Unlock()
	return err

}

func (a *Client) sendMessage(msg string, readTimeout time.Duration) error {
	err := a.writeMessage(msg)
	if err != nil {
		return err
	}
	isControlMessage := strings.HasPrefix(msg, CONTROL_HEADER)
	if isControlMessage && readTimeout < ReadWaitTimeout {
		readTimeout = ReadWaitTimeout
	}
	err = a.readServerResponse(readTimeout)
	if err != nil {
		a.logger.Warn("read failed", zap.Any("agentId", a.AgentID), zap.Error(err))
		return err
	}

	for a.CurrentRemoteFile != nil {
		err = a.readServerResponse(readTimeout)
		if err != nil {
			a.logger.Warn("read failed", zap.Any("agentId", a.AgentID), zap.Error(err))
			return err
		}
	}

	return nil
}

func (a *Client) readServerResponse(timeout time.Duration) error {
	if a.conn == nil {
		return fmt.Errorf("client is not connected")
	}

	var buf bytes.Buffer
	totallyRead := 0
	messagesRead := 0

	totallyRead, shouldReturn, err := readResponse(timeout, a, totallyRead, messagesRead, &buf)
	if err != nil {
		return err
	}
	if shouldReturn {
		return nil
	}
	if totallyRead == 0 {
		if a.CurrentRemoteFile != nil {
			a.cacheFileHash(a.CurrentRemoteFile.Filename, a.CurrentRemoteFile.Hash, a.CurrentRemoteFile.Content.String())
		}
		return nil
	}

	rawMsg := buf.Bytes()
	// fmt.Printf("\n%d -> '%s\n", len(rawMsg), string(rawMsg[:int(math.Min(float64(len(rawMsg)), 128))]))
	for {
		var msgSize uint32
		if !a.UDP {
			if len(rawMsg) < 4 {
				AgentCollector.MessageError(a, 1)
				return errors.New("message too short")
			}
			msgSize = binary.LittleEndian.Uint32(rawMsg)
			rawMsg = rawMsg[4:]
			totallyRead -= 4
		} else {
			msgSize = uint32(len(rawMsg))
		}

		if int(msgSize) > len(rawMsg) {
			AgentCollector.MessageError(a, 1)
			return errors.New("message exceeds buffer")
		}
		a.evtCount++
		msg, err := a.decryptMessage(rawMsg[:msgSize], msgSize)
		// fmt.Printf("%d\t%d\t'%s'\t%s\n", totallyRead, msgSize, msg, err)
		if err != nil {
			AgentCollector.MessageError(a, 1)
			return err
		}
		a.receivedCount++

		// parse result - get counters
		msg = msg[32:]
		globalCount, err := strconv.Atoi(msg[5:15])
		if err != nil {
			AgentCollector.MessageError(a, 1)
			return err
		}
		localCount, err := strconv.Atoi(msg[16:20])
		if err != nil {
			AgentCollector.MessageError(a, 1)
			return err
		}
		msg = msg[21:]
		localCountU := uint(localCount)
		globalCountU := uint(globalCount)
		if globalCountU == a.globalCount && (localCountU == a.localCount) {
			// normal status, nothing to report
		} else if globalCountU > a.globalCount || (globalCountU == a.globalCount && localCountU > a.localCount) {
			a.logger.Debug(fmt.Sprintf("Updated to remote counters %d:%d (%d:%d)", localCountU, globalCountU, a.localCount, a.globalCount), zap.Skip())
			// move one ahead
			localCountU++
		} else {
			a.logger.Info(fmt.Sprintf("Unexpected counter %d:%d (%d:%d)", localCountU, globalCountU, a.localCount, a.globalCount), zap.Skip())
		}
		a.localCount = localCountU
		a.globalCount = globalCountU

		a.WriteClientCounter()
		AgentCollector.MessagesReceived(a, 1)
		// rand1 := msg[:5]
		//fmt.Printf("packet-received: bytes=%d (%s:%d:%d) '%s'\n", nRead, rand1, globalCount, localCount, msg)
		// empty buffer for next read

		err = a.handleResponse(msg)
		if err != nil {
			return err
		}
		messagesRead++

		rawMsg = rawMsg[msgSize:]
		if len(rawMsg) == 0 {
			break
		}
	}
	return nil
}

func readResponse(timeout time.Duration, a *Client, totallyRead int, messagesRead int, buf *bytes.Buffer) (int, bool, error) {
	deadline := time.Now().Add(timeout)

	buffer := make([]byte, maxBufferSize)
	for {
		err := a.conn.SetReadDeadline(deadline)
		if err != nil {
			a.logger.Warn("setReadDeadlineFailed", zap.Any("agentId", a.AgentID), zap.Any("deadline", deadline), zap.Error(err))
			return 0, true, err
		}

		nRead, err := a.conn.Read(buffer)
		a.receivedBytes = a.receivedBytes + uint64(nRead)
		a.receivedBytesTotal = a.receivedBytesTotal + uint64(nRead)
		AgentCollector.BytesRead(a, nRead)

		if nRead == 0 {
			if oe, ok := err.(*net.OpError); ok && oe.Err != os.ErrDeadlineExceeded {
				return 0, true, err
			}
			break
		}
		if err != nil {
			if oe, ok := err.(*net.OpError); ok && totallyRead > 0 && oe.Err == os.ErrDeadlineExceeded {
				break
			} else {
				if messagesRead > 0 {
					return 0, true, nil
				}
				return 0, true, err
			}
		}
		buf.Write(buffer[:nRead])
		totallyRead += nRead
	}
	buffer = nil
	return totallyRead, false, nil
}

// Connect connect and do a handshake
func (a *Client) Connect(isStartup bool) error {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.sessionID = time.Now().UnixMicro()
	a.connected = false
	var err error
	// try to re-register agent if the connection has been closed before
	AgentCollector.Disconnect(a)
	AgentCollector.TryConnect(a)

	var localAddr net.Addr
	if a.UDP {
		a.logger.Debug("connect", zap.Any("agentId", a.AgentID), zap.String("protocol", "udp"), zap.String("server", a.Server))
		a.conn, err = net.Dial("udp", fmt.Sprintf("%s:%d", a.Server, a.Port))
		if err != nil {
			a.sessionID = 0
			return err
		}
		localAddr = a.conn.LocalAddr().(*net.UDPAddr)
	} else {
		a.logger.Debug("connect", zap.Any("agentId", a.AgentID), zap.String("protocol", "tcp"), zap.String("server", a.Server))
		a.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", a.Server, a.Port))
		if err != nil {
			a.sessionID = 0
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
	err = a.sendMessage(msg, ReadWaitTimeout)
	if err != nil {
		a.sessionID = 0
		return err
	}

	if isStartup {
		// log start startup
		var agentIP string
		if a.AgentAllowedIPs == "any" {
			agentIP = "any"
		} else {
			agentIP = a.AgentIP
		}
		msg = fmt.Sprintf("ossec: Agent started: '%s->%s'.", a.AgentID, agentIP)
		msg = fmt.Sprintf("%c:%s:%s", LOCALFILE_MQ, "ossec", msg)
		err = a.writeMessage(msg)
		if err != nil {
			a.sessionID = 0
			return err
		}

		time.Sleep(1 * time.Second)
		err = a.pingServer()
		if err != nil {
			a.sessionID = 0
			return err
		}
	}

	AgentCollector.Connect(a)
	a.connected = true
	return nil
}

// ItemBuilder creates a new item and returns a pointer to it.
// This is used when we load a segment of the queue from disk.
func itemBuilder() any {
	return &QueuePosting{}
}

// Helper function to dynamically register a type with Gob
func registerType(obj any) {
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem() // Dereference pointer type
	}
	gob.Register(obj)
	fmt.Printf("Registered type: %s\n", typ.Name())
}

func encodeData(data any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeData(data []byte, obj any) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(obj)
}

func (a *Client) openQueue(ctx context.Context) (chan *QueuePosting, *dque.DQue, error) {
	q, err := dque.NewOrOpen("event-queue", a.basePath, 500, itemBuilder)
	queuePath := a.basePath + "/event-queue"
	if err != nil && strings.HasPrefix(err.Error(), "unable to create queue segment in "+queuePath) {
		a.logger.Warn("drop corrupt queue", zap.String("path", queuePath), zap.Error(err))
		os.RemoveAll(queuePath)
		q, err = dque.NewOrOpen("event-queue", a.basePath, 500, itemBuilder)
	}
	if err != nil {
		return nil, nil, err
	}

	a.logger.Info("queue opened", zap.String("path", queuePath), zap.Error(err))

	input := make(chan *QueuePosting, 100)

	go func() {
		for {
			for msg := range input {
				if msg.Timestamp.Unix() == 0 {
					msg.Timestamp = time.Now()
				}

				if msg == nil {
					a.logger.Warn("enqueue", zap.Any("agentId", a.AgentID), zap.String("problem", "nil item"))
					continue
				}

				if err = q.Enqueue(msg); err == nil {
					AgentCollector.Enqueue(a)
				} else {
					if strings.Contains(err.Error(), "gob: type not registered") {
						a.logger.Info("enqueueItem - late register type", zap.Any("agentId", a.AgentID), zap.Any("item", msg), zap.String("cause", err.Error()))
						registerType(msg)
						if err = q.Enqueue(msg); err == nil {
							AgentCollector.Enqueue(a)
						}
					}
					if err != nil {
						a.logger.Error("enqueueItem", zap.Any("agentId", a.AgentID), zap.Any("item", msg), zap.Error(err))
						AgentCollector.SetQueueSize(a, q.Size())
					}
				}
			}
			if ctx.Err() != nil {
				break
			}
		}
	}()

	return input, q, err
}

// AgentLoop Process messages and keep track of connection status
func (a *Client) AgentLoop(ctx context.Context, closeOnError bool) (chan *QueuePosting, chan any, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	a.ctx = ctx

	// make the context cancable
	out := make(chan any, 10)
	a.outChannel = out
	var err error
	tries := 0
	for err = a.Connect(true); err != nil; {
		tries++
		if tries > 10 {
			return nil, out, err
		}
		a.logger.WithOptions(zap.WithCaller(false)).Warn("connect failed", zap.Any("agentId", a.AgentID), zap.Int("try", tries), zap.String("error", err.Error()))
		time.Sleep(time.Second * 2)
	}

	input, q, err := a.openQueue(ctx)
	if err != nil {
		return nil, out, err
	}
	AgentCollector.SetQueueSize(a, q.Size())

	go func() {
		defer func() {
			a.logger.Error("shutdown", zap.Any("agentId", a.AgentID))
			a.Close()
			q.Close()
		}()

		nextSysinfoUpdate := -1

		for {
			if a.CurrentRemoteFile != nil {
				a.logger.Debug("fileTransfer", zap.Any("agentId", a.AgentID), zap.String("fileName", a.CurrentRemoteFile.Filename))
			} else {
				pingWait := ratelimit.New(1) // per second

				for {
					if ctx.Err() != nil {
						out <- err
						break
					}

					select {
					case evt, ok := <-a.outChannel:
						if ok {
							a.logger.Info("eventRead", zap.Any("agentId", a.AgentID), zap.Any("event", evt))
							fmt.Printf("Value %v was read.\n", evt)
							// <- &FileUpdatedEvent{a.CurrentRemoteFile}
						}
					default:

					}

					loopEntry := time.Now()
					loopExit := loopEntry.Add(time.Second * (NotifyTime - 1))
					// once a second check if there is any message
					for {
						item, dqErr := q.Peek()
						if dqErr == dque.ErrEmpty {
							AgentCollector.SetQueueSize(a, 0)
							// a.logger.Debug("dequeue", zap.Any("agentId", a.AgentID), zap.String("problem", "queue empty"))
							break
						}
						if dqErr != nil {
							a.logger.Error("dequeue", zap.Any("agentId", a.AgentID), zap.Error(dqErr))
							AgentCollector.SetQueueSize(a, q.Size())
							break
						}

						if item == nil {
							a.logger.Warn("dequeue", zap.Any("agentId", a.AgentID), zap.String("problem", "nil item"))
							AgentCollector.SetQueueSize(a, q.Size())
							continue
						}

						if msg, ok := item.(*QueuePosting); ok {
							var b []byte
							var err error

							switch v := msg.Raw.(type) {
							case int:
								b = []byte(strconv.Itoa(v))
							case float64:
								b = []byte(fmt.Sprintf("%f", v))
							case string:
								b = []byte(v)
							default:
								b, err = json.Marshal(msg.Raw)
							}

							if err != nil {
								a.logger.Error("marshall", zap.Any("agentId", a.AgentID), zap.Error(err))
								item = nil
								q.Dequeue()
								AgentCollector.Dequeue(a)
								continue
							}

							if msg.TargetQueue == 0 {
								msg.TargetQueue = LOCALFILE_MQ
							}

							var wireMsg string
							if msg.TargetQueue == LOCALFILE_MQ {
								wireMsg = fmt.Sprintf("%c:%s:%s %s %s:%s", msg.TargetQueue, msg.Location, msg.Timestamp.UTC().Format("Jan 02 15:04:05"), a.AgentName, msg.ProgramName, string(b))
							} else if msg.TargetQueue == SECURE_MQ {
								wireMsg = fmt.Sprintf("%c:%s->%s", msg.TargetQueue, msg.Location, string(b))
							} else {
								wireMsg = fmt.Sprintf("%c:%s:%s", msg.TargetQueue, msg.Location, string(b))
							}

							err = a.WriteMessage(wireMsg)

							item = nil
							if err != nil {
								a.Close()
								break
							}
						} else {
							a.logger.Error("dequeue", zap.Any("agentId", a.AgentID), zap.Error(fmt.Errorf("invalid queue content")))
						}
						// remove last item from queue
						_, err = q.Dequeue()
						AgentCollector.Dequeue(a)
						if err != nil {
							a.logger.Error("dequeue", zap.Any("agentId", a.AgentID), zap.Error(err))
							AgentCollector.SetQueueSize(a, q.Size())
							item = nil
						}
						// make sure, we take a pause
						if time.Now().After(loopExit) {
							break
						}
					}

					// take a breath
					// prev := time.Now()
					//now :=
					pingWait.Take()
					// a.logger.Debug("pingWait", zap.Any("agentId", a.AgentID), zap.Duration("pingWait", now.Sub(prev)))

					if !a.IsConencted() {
						a.logger.Warn("disconnectedUnexpectedly", zap.Any("agentId", a.AgentID))
						break
					}
				}
			}

			if !a.IsConencted() {
				// abort any open download
				if a.CurrentRemoteFile != nil {
					a.CurrentRemoteFile = nil
				}
				tries := 0
				a.logger.Info("tryReconnect", zap.Any("agentId", a.AgentID))
				for err = a.Connect(false); err != nil; {
					time.Sleep(time.Second * 2)
					tries++
					if tries > 10 {
						a.logger.Panic("reconnectFailed", zap.Any("agentId", a.AgentID), zap.Error(err))
						break
					}
					a.logger.Info("tryReconnect", zap.Any("agentId", a.AgentID))
				}
			} else {
				err = a.PingServer()
				nextSysinfoUpdate--
				if nextSysinfoUpdate <= 0 {
					a.Scanner.PostSysinfo(input)
					nextSysinfoUpdate = SysinfoInterval
				}
			}
		}
	}()

	return input, out, nil
}
