package ossec

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
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
	PingIntervall = 60
)

// Client allowes to handshake with the server to reach a pending state (which allowes the agent to become a group member)
type Client struct {
	*AgentKey
	Server            string
	Port              uint16
	UDP               bool
	basePath          string
	remotePath        string
	localCount        uint
	globalCount       uint
	evtCount          uint
	cOrigSize         uint
	cCompSize         uint
	EncryptionMethod  EncryptionMethod
	ClientName        string
	ClientVersion     string
	ctx               context.Context
	conn              net.Conn
	mx                sync.Mutex
	logger            *zap.Logger
	connected         bool
	ratelimit         ratelimit.Limiter
	outChannel        chan interface{}
	RemoteFiles       map[string]RemoteFileInfo
	CurrentRemoteFile *RemoteFileInfo
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
	gob.Register(map[string]interface{}{})
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
		/**		if encryptionMethod == EncryptionMethodAES {
			return errors.New("AES is currently not supported")
		}
		*/
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
		EncryptionMethod: EncryptionMethodAES,
		ClientName:       "go-wazuh",
		ClientVersion:    "v1.0.0",
		ratelimit:        ratelimit.New(SendRateLimit), // per second
		RemoteFiles:      make(map[string]RemoteFileInfo),
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

	if a.basePath == "" {
		if _, err := os.Stat("/var/tmp"); !os.IsNotExist(err) {
			a.basePath = "/var/tmp"
		} else if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
			a.basePath = dir
		} else {
			a.basePath = os.TempDir()
		}
	}

	a.remotePath = filepath.Join(a.basePath, "remote")
	if _, err := os.Stat(a.remotePath); os.IsNotExist(err) {
		os.MkdirAll(a.remotePath, 0770)
	}
	return a, nil
}

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

	LOCALFILE_MQ = '1'
	SYSLOG_MQ    = '2'
	HOSTINFO_MQ  = '3'
	SECURE_MQ    = '4'
	SYSCHECK_MQ  = '8'
	ROOTCHECK_MQ = '9'

	maxBufferSize        = 1024 * 60
	ReadWaitTimeout      = time.Duration(30 * time.Second)
	ReadImmediateTimeout = time.Duration(1 * time.Second)
)

func (a *Client) IsConencted() bool {
	return a.connected
}

func (a *Client) close(sendCloseMsg bool) error {
	if a.connected {
		if sendCloseMsg {
			msg := fmt.Sprintf("ossec: ossec: Agent disconnected: '%s-%s'.", a.AgentName, a.AgentIP)
			msg = fmt.Sprintf("%c:%s:%s", LOCALFILE_MQ, "ossec", msg)
			err := a.writeMessage(msg)
			if err != nil {
				a.connected = false
				a.conn.Close()
				return err
			}
		}
		a.connected = false
		return a.conn.Close()
	}
	return nil
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (a *Client) Close() error {
	return a.close(true)
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

func (a *Client) handleResponse(response string) error {

	if strings.HasPrefix(string(response), CONTROL_HEADER) {
		if strings.HasPrefix(string(response), FILE_UPDATE_HEADER) {
			fieleSpecs := strings.Split(strings.Trim(response[11:], "\n \t"), " ")
			if len(fieleSpecs) == 2 {
				if existingFile, ok := a.RemoteFiles[fieleSpecs[1]]; ok && existingFile.Hash == fieleSpecs[0] {
					return nil
				}
				if a.logger != nil {
					a.logger.Debug("receive file", zap.String("fileName", fieleSpecs[1]))
				}
				a.CurrentRemoteFile = &RemoteFileInfo{
					Filename: fieleSpecs[1],
					Hash:     fieleSpecs[0],
					Content:  bytes.NewBuffer(nil),
				}
			}
			return nil
		}
		if strings.HasPrefix(string(response), FILE_CLOSE_HEADER) {
			if a.logger != nil {
				a.logger.Debug("done receive file", zap.Int("len", a.CurrentRemoteFile.Content.Len()))
			}
			a.cacheFileHash(a.CurrentRemoteFile.Filename, a.CurrentRemoteFile.Hash, a.CurrentRemoteFile.Content.String())
			a.outChannel <- &FileUpdatedEvent{a.CurrentRemoteFile}
			return nil
		}
		if string(response) == HC_ACK {
			return nil
		}
		a.logger.Info("control message", zap.Any("agentId", a.AgentID), zap.String("msg", string(response)))

	} else {
		if a.CurrentRemoteFile != nil {
			a.CurrentRemoteFile.Content.WriteString(response)
			a.cacheFileHash(a.CurrentRemoteFile.Filename, a.CurrentRemoteFile.Hash, a.CurrentRemoteFile.Content.String())

			/*if a.logger != nil {
				a.logger.Debug("receivedFilecontent", zap.Any("agentId", a.AgentID), zap.String("filename", a.CurrentRemoteFile.Filename), zap.Int("read", len(response)), zap.Int("total", a.CurrentRemoteFile.Content.Len()), zap.String("msg", string(response)))
			}
			*/
			return nil
		}
		if a.logger != nil {
			a.logger.Debug("receivedMessage", zap.Any("agentId", a.AgentID), zap.String("msg", string(response)))
		}
	}
	return nil
}

func (a *Client) getSharedFiles() string {
	b := strings.Builder{}
	searchPath := filepath.Join(a.remotePath, "*")
	files, err := filepath.Glob(searchPath)
	if err == nil && len(files) > 0 {
		for _, filename := range files {
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				a.logger.Warn("corrupt file", zap.String("filename", filename), zap.Error(err))
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
		b.WriteString("00000000000000000000000000000000 merged.mg\n")
	}
	return b.String()
}

func (a *Client) cacheFileHash(fileName string, hash string, content string) error {
	fullpath := filepath.Join(a.remotePath, fileName)
	err := ioutil.WriteFile(fullpath, []byte(content), 0644)
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
	labelWodle := fmt.Sprintf("#\"_opserve_wodle\":\"%s\"", a.ClientName)
	agentConfigMd5 := fmt.Sprintf("%00x", md5.Sum([]byte(a.AgentName)))
	un := goInfo.GetInfo()

	osInfo := sysinfo.GetOSInfo()

	osRel := osInfo.Name
	if strings.HasPrefix(strings.ToLower(osRel), strings.ToLower(osInfo.Vendor+" ")) {
		osRel = osRel[len(osInfo.Vendor)+1:]
	}
	aname := fmt.Sprintf("%s |%s |%s |%s |%s [%s|%s: %s] - %s %s",
		firstUpper(un.OS),
		un.Hostname,
		un.Core,
		un.Kernel,
		un.Platform,
		firstUpper(runtime.GOOS),
		runtime.GOARCH,
		osRel,
		a.ClientName, a.ClientVersion)

	labels := ""

	sharedFiles := a.getSharedFiles()

	// test example:
	//"Linux |debian10 |4.19.0-9-amd64 |#1 SMP Debian 4.19.118-2+deb10u1 (2020-06-07) |x86_64 \
	// [Debian GNU/Linux|debian: 10 (buster)] - Wazuh v3.13.0 / ab73af41699f13fdd81903b5f23d8d00\nfd756ba04d9c32c8848d4608bec41251 \
	// merged.mg\n#\"_agent_ip\":192.168.0.143\n"
	msg := fmt.Sprintf("%s%s / %s\n%s%s%s\n%s\n", CONTROL_HEADER, aname, agentConfigMd5, labels, sharedFiles, labelIP, labelWodle)

	err := a.sendMessage(msg, ReadWaitTimeout)
	if err != nil {
		return err
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

	// ensure ratelimit is honored
	prev := time.Now()
	now := a.ratelimit.Take()

	ret, err := a.conn.Write(encryptedMsg)
	if err != nil {
		if a.logger != nil {
			a.logger.Debug("writeMessage", zap.Any("agentId", a.AgentID), zap.String("msg", msg), zap.Int("result", ret), zap.Duration("rateWait", now.Sub(prev)), zap.Error(err))
		}
		a.close(false)
		return err
	}
	if a.logger != nil {
		a.logger.Debug("writeMessage", zap.Any("agentId", a.AgentID), zap.String("msg", msg), zap.Int("result", ret), zap.Duration("rateWait", now.Sub(prev)))
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
	return a.readServerResponse(readTimeout)

}

func (a *Client) readServerResponse(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	err := a.conn.SetReadDeadline(deadline)
	if err != nil {
		println(err.Error())
		return err
	}

	var buf bytes.Buffer
	buffer := make([]byte, maxBufferSize)
	totallyRead := 0
	messagesRead := 0
	for {
		for {
			nRead, err := a.conn.Read(buffer)
			if nRead == 0 {
				break
			}
			if err != nil {
				if oe, ok := err.(*net.OpError); ok && totallyRead > 0 && oe.Err == os.ErrDeadlineExceeded {
					break
				} else {
					if messagesRead > 0 {
						// time out after at least a message == EOF... not a real error
						return nil
					}
					// time out without any message... not expected
					// fmt.Printf("%d\t%d\t'%s'\t%s\n", totallyRead, -1, "<nil>", err)
					return err
				}
			}

			buf.Write(buffer[:nRead])
			totallyRead += nRead
			if nRead < maxBufferSize {
				break
			}
		}
		if totallyRead == 0 {
			return nil
		}
		rawMsg := buf.Bytes()
		// fmt.Printf("\n%d -> '%s\n", len(rawMsg), string(rawMsg[:int(math.Min(float64(len(rawMsg)), 128))]))

		var msgSize uint32
		if !a.UDP {
			if len(rawMsg) < 4 {
				return errors.New("message too short")
			}
			msgSize = binary.LittleEndian.Uint32(rawMsg)
			rawMsg = rawMsg[4:]
			totallyRead -= 4
		} else {
			msgSize = uint32(len(rawMsg))
		}

		if int(msgSize) > len(rawMsg) {
			return errors.New("message exceeds buffer")
		}
		msg, err := a.decryptMessage(rawMsg[:msgSize], msgSize)
		// fmt.Printf("%d\t%d\t'%s'\t%s\n", totallyRead, msgSize, msg, err)
		if err != nil {
			return err
		}

		// parse result - get counters
		msg = msg[32:]
		globalCount, err := strconv.Atoi(msg[5:15])
		if err != nil {
			return err
		}
		localCount, err := strconv.Atoi(msg[16:20])
		if err != nil {
			return err
		}
		msg = msg[21:]
		a.localCount = uint(localCount)
		a.globalCount = uint(globalCount)
		// rand1 := msg[:5]
		// fmt.Printf("packet-received: bytes=%d (%s:%d:%d) '%s'\n", nRead, rand1, globalCount, localCount, msg)
		// empty buffer for next read
		buf.Reset()
		// already read a port of the next message?
		if totallyRead > int(msgSize) {
			totallyRead -= int(msgSize)
			buf.Write(rawMsg[int(msgSize):])
			totallyRead = 0
		} else {
			totallyRead = 0
			buffer = make([]byte, maxBufferSize)
		}

		err = a.handleResponse(msg)
		if err != nil {
			return err
		}
		messagesRead++
		deadline := time.Now().Add(ReadImmediateTimeout)
		err = a.conn.SetReadDeadline(deadline)
		if err != nil {
			println(err.Error())
			return err
		}
	}
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
	err = a.sendMessage(msg, ReadWaitTimeout)
	if err != nil {
		return err
	}

	if isStartup {
		// log start startup
		msg = fmt.Sprintf("ossec: Agent started: '%s->%s'.", a.AgentID, a.AgentIP)
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

// ItemBuilder creates a new item and returns a pointer to it.
// This is used when we load a segment of the queue from disk.
func itemBuilder() interface{} {
	return &QueuePosting{}
}

func (a *Client) openQueue(ctx context.Context) (chan *QueuePosting, *dque.DQue, error) {

	q, err := dque.NewOrOpen("event-queue", a.basePath, 500, itemBuilder)
	if err != nil {
		return nil, nil, err
	}

	input := make(chan *QueuePosting, 100)

	go func() {
		for {
			for msg := range input {
				if msg.Timestamp == time.Unix(0, 0) {
					msg.Timestamp = time.Now()
				}
				if err = q.Enqueue(msg); err != nil {
					a.logger.Error("enqueue item", zap.Any("item", msg), zap.Error(err))
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
func (a *Client) AgentLoop(ctx context.Context, closeOnError bool) (chan *QueuePosting, chan interface{}, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	a.ctx = ctx

	// make the context cancable
	out := make(chan interface{})
	a.outChannel = out
	var err error
	for err = a.Connect(true); err != nil; {
		a.logger.WithOptions(zap.WithCaller(false)).Warn("connect failed", zap.Any("agentId", a.AgentID), zap.String("error", err.Error()))
		time.Sleep(time.Second * 10)
	}

	input, q, err := a.openQueue(ctx)
	if err != nil {
		return nil, out, err
	}

	go func() {
		defer a.Close()
		defer q.Close()

		for {
			loopEntry := time.Now()
			for t := 0; t < PingIntervall; t++ {
				if ctx.Err() != nil {
					out <- err
					break
				}

				// once a second check if there is any message
				for item, dqErr := q.Peek(); dqErr != dque.ErrEmpty && item != nil; {
					if dqErr != nil {
						a.logger.Error("dequeue", zap.Error(dqErr))
						break
					}

					if msg, ok := item.(*QueuePosting); ok {
						b, err := json.Marshal(msg.Raw)
						if err != nil {
							a.logger.Error("marshall", zap.Error(err))
							item = nil
							q.Dequeue()
							continue
						}

						if msg.TargetQueue == 0 {
							msg.TargetQueue = LOCALFILE_MQ
						}

						wireMsg := fmt.Sprintf("%c:%s:%s %s %s:%s", msg.TargetQueue, msg.Location, msg.Timestamp.UTC().Format("Jan 02 15:04:05"), a.AgentName, msg.ProgramName, string(b))
						err = a.WriteMessage(wireMsg)
						item = nil
						if err != nil {
							break
						}
					} else {
						a.logger.Error("dequeue", zap.Error(fmt.Errorf("invalid queue content")))
					}
					// remove last item from queue
					q.Dequeue()
					// make sure, we take a pause
					if loopEntry.Add(time.Second*PingIntervall - 1).After(time.Now()) {
						break
					}
				}
				// take a breath
				time.Sleep(time.Second * 1)
				if !a.IsConencted() {
					a.logger.Warn("disconnected")
					break
				}
			}
			if !a.IsConencted() {
				a.logger.Debug("try reconnect", zap.Any("agentId", a.AgentID))
				for err = a.Connect(false); err != nil; {
					time.Sleep(time.Second * 10)
					a.logger.Debug("try reconnect", zap.Any("agentId", a.AgentID))
				}
			} else {
				err = a.PingServer()
			}

		}
	}()

	return input, out, nil
}
