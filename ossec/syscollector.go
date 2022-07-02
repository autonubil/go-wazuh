package ossec

// See: https://documentation.wazuh.com/current/user-manual/capabilities/syscollector.html
// See: https://github.com/wazuh/wazuh/blob/master/src/analysisd/decoders/syscollector.c
// See: https://github.com/wazuh/wazuh/blob/master/framework/wazuh/core/syscollector.py

import (
	"encoding/gob"
	"net"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	// https://www.socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

type SysCollector struct {
	agent      *Client
	scanId     int64
	localPorts map[string]*PortInfo
	scanTime   string
	mxScan     sync.Mutex
}

var Scanner *SysCollector

func NewScanner(client *Client) *SysCollector {
	return &SysCollector{
		scanId:     int64(time.Now().Unix()),
		agent:      client,
		localPorts: make(map[string]*PortInfo),
	}
}

func init() {
	gob.Register(Sysinfo{})
	gob.Register(OS{})
	gob.Register(Hardware{})
	gob.Register(Package{})
	gob.Register(Network{})
	gob.Register(Process{})
	gob.Register(Port{})
}

const (
	TYPE_OS          = "OS"
	TYPE_HARDWARE    = "hardware"
	TYPE_NETWORK     = "network"
	TYPE_NETWORK_END = "network_end"
	TYPE_PROCESS     = "process"
	TYPE_PROCESS_END = "process_end"
	TYPE_PORT        = "port"
	TYPE_PORT_END    = "port_end"
	TYPE_PACKAGE     = "program"
	TYPE_PACKAGE_END = "program_end"
)

func (s *SysCollector) PostSysinfo(input chan *QueuePosting) {
	s.mxScan.Lock()
	defer s.mxScan.Unlock()
	s.scanId++
	s.scanTime = getScanTime()

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         s.NewOS(),
	}

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         s.NewHardware(),
	}

	s.sendNetworks(input)
	s.sendPorts(input)
	s.sendPackages(input)

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         s.NewProcess(TYPE_PROCESS),
	}

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         &Process{Sysinfo: s.NewSysinfo(TYPE_PROCESS_END)},
	}

}

func (s *SysCollector) sendPackages(input chan *QueuePosting) {
	pwd, _ := os.Getwd()
	for _, searchPath := range []string{"/", pwd, s.agent.GetBasePath()} {
		filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if matched, err := filepath.Match("*.sbom", filepath.Base(path)); err != nil {
				return err
			} else if matched {
				s.sendSBOM(input, path)
			}
			return nil
		})
	}
}

func (s *SysCollector) sendSBOM(input chan *QueuePosting, path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	bom := new(cdx.BOM)
	decoder := cdx.NewBOMDecoder(file, cdx.BOMFileFormatJSON)
	if err = decoder.Decode(bom); err != nil {
		return
	}

	for _, component := range *bom.Components {
		pkg := s.NewPackageFromComponent(component)
		if pkg != nil {
			input <- &QueuePosting{
				Location:    SYSCOLLECTOR_MOD,
				TargetQueue: SYSCOLLECTOR_MQ,
				Timestamp:   time.Now(),
				Raw:         pkg,
			}
		}
	}

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         &Package{Sysinfo: s.NewSysinfo(TYPE_PACKAGE_END)},
	}

}

func (s *SysCollector) sendPorts(input chan *QueuePosting) {
	for _, pi := range s.localPorts {
		input <- &QueuePosting{
			Location:    SYSCOLLECTOR_MOD,
			TargetQueue: SYSCOLLECTOR_MQ,
			Timestamp:   time.Now(),
			Raw:         &Port{Sysinfo: s.NewSysinfo(TYPE_PORT), PortInfo: s.NewPortInfo(pi)},
		}
	}

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         &Port{Sysinfo: s.NewSysinfo(TYPE_PORT_END)},
	}
}

func (s *SysCollector) sendNetworks(input chan *QueuePosting) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, iface := range ifaces {
		net, isLocal := s.NewNetwork(iface)
		if !isLocal {
			input <- &QueuePosting{
				Location:    SYSCOLLECTOR_MOD,
				TargetQueue: SYSCOLLECTOR_MQ,
				Timestamp:   time.Now(),
				Raw:         net,
			}
		}
	}
	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         &Network{Sysinfo: s.NewSysinfo(TYPE_NETWORK_END)},
	}

}

type Sysinfo struct {
	// ScanTime  string     `json:"scan_time,omitempty"`
	Type     string  `json:"type"`
	Checksum *string `json:"checksum,omitempty"`
	ID       int64   `json:"ID,omitempty"`
	ItemID   *int64  `json:"item_id,omitempty"`
	ScanTime string  `json:"timestamp"`
}

func getScanTime() string {
	return time.Now().Format("2006/01/02 03:04:05")
}

func (s SysCollector) NewSysinfo(typ string) *Sysinfo {
	return &Sysinfo{
		Type:     typ,
		ID:       s.scanId,
		ScanTime: s.scanTime,
	}
}

type HardwareInventory struct {
	BoardSerial *string  `json:"board_serial,omitempty"`
	CPUName     *string  `json:"cpu_name,omitempty"`
	CPUCores    *int     `json:"cpu_cores,omitempty"`
	CPUMhz      *float64 `json:"cpu_mhz,omitempty"`
	RamFree     *uint64  `json:"ram_free,omitempty"`
	RamTotal    *uint64  `json:"ram_total,omitempty"`
	RamUsage    *float64 `json:"ram_usage,omitempty"`
}

type Hardware struct {
	// R"({"board_serial":"Intel Corporation","scan_time":"2020/12/28 21:49:50", "cpu_MHz":2904,"cpu_cores":2,"cpu_name":"Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz", "ram_free":2257872,"ram_total":4972208,"ram_usage":54})")));
	*Sysinfo
	Inventory HardwareInventory `json:"inventory"`
}

func (s *SysCollector) NewHardware() *Hardware {
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	ramFree := vmStat.Free / 1024
	ramTotal := vmStat.Total / 1024
	ramUsage := vmStat.UsedPercent / 1024

	numCpu := runtime.NumCPU()
	hw := &Hardware{
		Sysinfo: s.NewSysinfo(TYPE_HARDWARE),
		Inventory: HardwareInventory{
			CPUName:  &cpuStat[0].ModelName,
			CPUCores: &numCpu,
			CPUMhz:   &cpuStat[0].Mhz,
			RamFree:  &ramFree,
			RamTotal: &ramTotal,
			RamUsage: &ramUsage,
		},
	}
	return hw
}

type OSInventory struct {
	OSName       *string `json:"os_name,omitempty"`
	OSVersion    *string `json:"os_version,omitempty"`
	OSCodename   *string `json:"os_codename,omitempty"`
	OSMajor      *string `json:"os_major,omitempty"`
	OSMinor      *string `json:"os_minor,omitempty"`
	OSBuild      *string `json:"os_build,omitempty"`
	Hostname     *string `json:"hostname,omitempty"`
	OSRelease    *string `json:"os_release,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
}
type OS struct {
	//   R"({"architecture":"x86_64","scan_time":"2020/12/28 21:49:50", "hostname":"UBUNTU","os_build":"7601","os_major":"6","os_minor":"1","os_name":"Microsoft Windows 7","os_release":"sp1","os_version":"6.1.7601"})")));
	*Sysinfo
	Inventory OSInventory `json:"inventory"`
}

func (s *SysCollector) NewOS() *OS {
	arch := runtime.GOARCH
	os := &OS{
		Sysinfo: s.NewSysinfo(TYPE_OS),
		Inventory: OSInventory{
			OSName:       &s.agent.osInfo.Name,
			OSVersion:    &s.agent.osInfo.Version,
			Hostname:     &s.agent.un.Hostname,
			OSRelease:    &s.agent.osInfo.Release,
			Architecture: &arch,
		},
	}
	return os
}

// R"([{"architecture":"amd64","scan_time":"2020/12/28 21:49:50", "group":"x11","name":"xserver-xorg","priority":"optional","size":"411","source":"xorg","version":"1:7.7+19ubuntu14","os_patch":""},{"hotfix":"KB4586786"}])")));

type ProcessEntry struct {
	Name      *string  `json:"name,omitempty"`
	Cmd       *string  `json:"cmd,omitempty"`
	ArgVs     []string `json:"argvs,omitempty"`
	EGroup    *string  `json:"egroup,omitempty"`
	EUser     *string  `json:"euser,omitempty"`
	FGroup    *string  `json:"fgroup,omitempty"`
	RGroup    *string  `json:"rgroup,omitempty"`
	RUser     *string  `json:"ruser,omitempty"`
	SGroup    *int     `json:"sgroup,omitempty"`
	SUser     *int     `json:"suser,omitempty"`
	State     *string  `json:"state,omitempty"`
	Nice      *int     `json:"nice,omitempty"`
	NLWP      *int     `json:"nlwp,omitempty"`
	PGrp      *int     `json:"pgrp,omitempty"`
	PID       *int     `json:"pid,omitempty"`
	PPID      *int     `json:"ppid,omitempty"`
	Priority  int      `json:"priority"`
	Processor *int     `json:"processor,omitempty"`
	Resident  *int     `json:"resident,omitempty"`
	Session   *int     `json:"session,omitempty"`
	Share     *int     `json:"share,omitempty"`
	Size      *uint64  `json:"size,omitempty"`
	STime     *int     `json:"stime,omitempty"`
	TGID      *int     `json:"tgid,omitempty"`
	TTY       *int     `json:"tty,omitempty"`
	UTime     *int     `json:"utime,omitempty"`
	VMSize    *uint64  `json:"vm_size,omitempty"`
}

type Process struct {
	// R"([{"egroup":"root","euser":"root","fgroup":"root","name":"kworker/u256:2-","scan_time":"2020/12/28 21:49:50", "nice":0,"nlwp":1,"pgrp":0,"pid":431625,"ppid":2,"priority":20,"processor":1,"resident":0,"rgroup":"root","ruser":"root","session":0,"sgroup":"root","share":0,"size":0,"start_time":9302261,"state":"I","stime":3,"suser":"root","tgid":431625,"tty":0,"utime":0,"vm_size":0}])")));

	// CREATE TABLE sys_processes ( scan_id INTEGER,
	//	scan_time TEXT, pid TEXT, name TEXT, state TEXT, ppid INTEGER, utime INTEGER, stime INTEGER, cmd TEXT, argvs TEXT, euser TEXT, ruser TEXT, suser TEXT, egroup TEXT, rgroup TEXT, sgroup TEXT, fgroup TEXT, priority INTEGER, nice INTEGER, size INTEGER, vm_size INTEGER, resident INTEGER, share INTEGER, start_time INTEGER, pgrp INTEGER, session INTEGER, nlwp INTEGER, tgid INTEGER, tty INTEGER, processor INTEGER,
	//	checksum TEXT NOT NULL CHECK (checksum <> ''), PRIMARY KEY (scan_id, pid));

	*Sysinfo
	ProcessDetails *ProcessEntry `json:"process,omitempty"`
}

func (s *SysCollector) NewProcess(typ string) *Process {
	var euser string
	var egroup string
	name := filepath.Base(os.Args[0])
	cmd := os.Args[0]
	pid := os.Getpid()
	ppid := os.Getppid()

	uid := os.Geteuid()
	euser = strconv.Itoa(uid)
	eUser, err := user.LookupId(euser)
	if err == nil {
		euser = eUser.Name
	}

	egid := os.Getegid()
	egroup = strconv.Itoa(egid)
	eGroup, err := user.LookupGroupId(egroup)
	if err == nil {
		egroup = eGroup.Name
	}
	state := "R"
	var mStats runtime.MemStats
	runtime.ReadMemStats(&mStats)

	p := &Process{
		Sysinfo: s.NewSysinfo(typ),
		ProcessDetails: &ProcessEntry{
			Name:     &name,
			Cmd:      &cmd,
			ArgVs:    os.Args[1:],
			EGroup:   &egroup,
			EUser:    &euser,
			PID:      &pid,
			Session:  &pid,
			PPID:     &ppid,
			State:    &state,
			Size:     &mStats.Alloc,
			VMSize:   &mStats.HeapAlloc,
			Priority: 0,
		},
	}
	return p
}

// R"([{"architecture":"amd64","scan_time":"2020/12/28 21:49:50", "group":"x11","name":"xserver-xorg","priority":"optional","size":"411","source":"xorg","version":"1:7.7+19ubuntu14","os_patch":""},{"hotfix":"KB4586786"}])")));
type PackageDetails struct {
	Format       *string `json:"format,omitempty"`
	Name         *string `json:"name,omitempty"`
	Priority     *string `json:"priority,omitempty"`
	Group        *string `json:"group,omitempty"`
	Size         *int64  `json:"size,omitempty"`
	Vendor       *string `json:"vendor,omitempty"`
	Version      *string `json:"version,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	MultiArch    *string `json:"multi-arch,omitempty"`
	Source       *string `json:"source,omitempty"`
	Description  *string `json:"description,omitempty"`
	InstallTime  *string `json:"install_time,omitempty"`
	Location     *string `json:"location,omitempty"`

	Triaged *string `json:"triaged,omitempty"` // read only
	CPE     *string `json:"cpe,omitempty"`     // read only
}
type Package struct {
	// R"({"iface":[{"address":"127.0.0.1","scan_time":"2020/12/28 21:49:50", "mac":"d4:5d:64:51:07:5d", "gateway":"192.168.0.1|600","broadcast":"127.255.255.255", "name":"ens1", "mtu":1500, "name":"enp4s0", "adapter":" ", "type":"ethernet", "state":"up", "dhcp":"disabled","iface":"Loopback Pseudo-Interface 1","metric":"75","netmask":"255.0.0.0","proto":"IPv4","rx_bytes":0,"rx_dropped":0,"rx_errors":0,"rx_packets":0,"tx_bytes":0,"tx_dropped":0,"tx_errors":0,"tx_packets":0, "IPv4":[{"address":"192.168.153.1","broadcast":"192.168.153.255","dhcp":"unknown","metric":" ","netmask":"255.255.255.0"}], "IPv6":[{"address":"fe80::250:56ff:fec0:8","dhcp":"unknown","metric":" ","netmask":"ffff:ffff:ffff:ffff::"}]}]})")));
	*Sysinfo
	Package *PackageDetails `json:"program"`
}

type NetworkInterface struct {
	Name      *string        `json:"name,omitempty"`
	Adapter   *string        `json:"adapter,omitempty"`
	Type      *string        `json:"type,omitempty"`
	State     *string        `json:"state,omitempty"`
	MAC       *string        `json:"mac,omitempty"`
	TXPackets *int64         `json:"tx_packets,omitempty"`
	RXPackets *int64         `json:"rx_packets,omitempty"`
	TXBytes   *int64         `json:"tx_bytes,omitempty"`
	RXBytes   *int64         `json:"rx_bytes,omitempty"`
	TXErrors  *int64         `json:"tx_errors,omitempty"`
	RXErrors  *int64         `json:"rx_errors,omitempty"`
	TXDropped *int64         `json:"tx_dropped,omitempty"`
	RXDropped *int64         `json:"rx_dropped,omitempty"`
	MTU       *int           `json:"mtu,omitempty"`
	IPv4      *IPAddressInfo `json:"IPv4,omitempty"`
	IPv6      *IPAddressInfo `json:"IPv6,omitempty"`
}

type IPv4Address struct {
}
type IPAddressInfo struct {
	Address   []string `json:"address,omitempty"`
	Netmask   []string `json:"netmask,omitempty"`
	Broadcast []string `json:"broadcast,omitempty"`
	Gateway   *string  `json:"gateway,omitempty"`
	// DHCP can be 'enabled', 'disabled', 'unknown', 'BOOTP'
	DHCP   string `json:"dhcp,omitempty"`
	Metric *int64 `json:"metric,omitempty"`
}

type Network struct {
	// R"({"iface":[{"address":"127.0.0.1","scan_time":"2020/12/28 21:49:50", "mac":"d4:5d:64:51:07:5d", "gateway":"192.168.0.1|600","broadcast":"127.255.255.255", "name":"ens1", "mtu":1500, "name":"enp4s0", "adapter":" ", "type":"ethernet", "state":"up", "dhcp":"disabled","iface":"Loopback Pseudo-Interface 1","metric":"75","netmask":"255.0.0.0","proto":"IPv4","rx_bytes":0,"rx_dropped":0,"rx_errors":0,"rx_packets":0,"tx_bytes":0,"tx_dropped":0,"tx_errors":0,"tx_packets":0, "IPv4":[{"address":"192.168.153.1","broadcast":"192.168.153.255","dhcp":"unknown","metric":" ","netmask":"255.255.255.0"}], "IPv6":[{"address":"fe80::250:56ff:fec0:8","dhcp":"unknown","metric":" ","netmask":"ffff:ffff:ffff:ffff::"}]}]})")));
	*Sysinfo
	Interface *NetworkInterface `json:"iface,omitempty"`
}

func (s *SysCollector) NewNetwork(intf net.Interface) (*Network, bool) {
	isLocal := false
	eth := "ethernet"
	state := "down"
	if intf.Flags|net.FlagUp == intf.Flags {
		state = "up"
	}
	network := &Network{
		Sysinfo: s.NewSysinfo(TYPE_NETWORK),
		Interface: &NetworkInterface{
			Name:  &intf.Name,
			MTU:   &intf.MTU,
			Type:  &eth,
			State: &state,
		},
	}
	addrs, err := intf.Addrs()
	if err == nil {
		ipv4 := &IPAddressInfo{DHCP: "unknown"}
		ipv6 := &IPAddressInfo{DHCP: "unknown"}
		for _, a := range addrs {
			if ipa, ok := a.(*net.IPNet); ok {
				if len(ipa.Mask) == 4 {
					ipv4.Address = append(ipv4.Address, ipa.IP.String())
					isLocal = isLocal || (strings.HasPrefix(ipa.IP.String(), "127.0.0."))
					ipv4.Netmask = append(ipv4.Netmask, net.IP(ipa.Mask).String())
					ipv4.Broadcast = append(ipv4.Broadcast, "")
				} else {
					ipv6.Address = append(ipv6.Address, ipa.IP.String())
					ipv6.Netmask = append(ipv6.Netmask, net.IP(ipa.Mask).String())
					ipv6.Broadcast = append(ipv6.Broadcast, "")
				}
			}
		}
		if len(ipv4.Address) > 0 {
			network.Interface.IPv4 = ipv4
		}
		if len(ipv6.Address) > 0 {
			network.Interface.IPv6 = ipv6
		}
	}

	if len(intf.HardwareAddr) > 0 {
		hwaddr := intf.HardwareAddr.String()
		network.Interface.MAC = &hwaddr
	}
	return network, isLocal
}

type PortInfo struct {
	Protocol   string  `json:"protocol,omitempty"`
	LocalIP    *string `json:"local_ip,omitempty"`
	RemoteIP   *string `json:"remote_ip,omitempty"`
	State      *string `json:"state,omitempty"`
	Pid        *int    `json:"PID,omitempty"`
	Process    *string `json:"process,omitempty"`
	LocalPort  *uint16 `json:"local_port,omitempty"`
	RemotePort *uint16 `json:"remote_port,omitempty"`
	TXQueue    *uint   `json:"tx_queue,omitempty"`
	RXQueue    *uint   `json:"rx_queue,omitempty"`
	Inode      *uint   `json:"inode,omitempty"`
}

type Port struct {
	// R"({"ports":[{"inode":0,"local_ip":"127.0.0.1","scan_time":"2020/12/28 21:49:50", "local_port":631,"pid":0,"process_name":"System Idle Process","protocol":"tcp","remote_ip":"0.0.0.0","remote_port":0,"rx_queue":0,"state":"listening","tx_queue":0}]})")));
	*Sysinfo
	PortInfo *PortInfo `json:"port,omitempty"`
}

func (s *SysCollector) NewPort(pi *PortInfo) *Port {
	cmd := os.Args[0]
	pid := os.Getpid()

	pi.Pid = &pid
	pi.Process = &cmd

	port := &Port{
		Sysinfo:  s.NewSysinfo(TYPE_PORT),
		PortInfo: pi,
	}
	return port
}

func (s *SysCollector) NewPortInfo(pi *PortInfo) *PortInfo {
	cmd := os.Args[0]
	pid := os.Getpid()

	pi.Pid = &pid
	pi.Process = &cmd

	return pi
}

func getPackage(component cdx.Component) *string {
	//
	if component.Type == cdx.ComponentTypeLibrary {
		if strings.HasPrefix(component.BOMRef, "pkg:") {
			typ := strings.Split(component.BOMRef[4:], "/")[0]
			switch typ {
			case "pacman":
				return &typ
			case "deb":
				return &typ
			case "rpm":
				return &typ
			case "win":
				return &typ
			case "pkg":
				return &typ
			default:
				//				fmt.Println(typ)
				//				return &typ
			}
		}

	}
	return nil
}

func (s *SysCollector) NewPackageFromComponent(component cdx.Component) *Package {
	arch := runtime.GOARCH
	// 0|2022/02/15 19:43:23|deb|libgssapi-krb5-2|optional|libs|426|Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>||1.17-6ubuntu4.1|amd64|same|krb5|MIT Kerberos runtime libraries - krb5 GSS-API Mechanism||1|||ca498815b5cf037988ce8a1e9ffe8183ae83e4f9|00d7743f300b6ade5f6eb75b2667ce30043b69a6
	//  scan_id INTEGER,    scan_time TEXT,    format TEXT NOT NULL CHECK (format IN ('pacman', 'deb', 'rpm', 'win', 'pkg')),    name TEXT,    priority TEXT,    section TEXT,    size INTEGER CHECK (size >= 0),    vendor TEXT,    install_time TEXT,    version TEXT,    architecture TEXT,    multiarch TEXT,    source TEXT,    description TEXT,    location TEXT,    triaged INTEGER(1),    cpe TEXT,    msu_name TEXT,    checksum TEXT NOT NULL CHECK (checksum <> ''),

	// fmt.Printf("%v\n", component)
	pkg := getPackage(component)
	if pkg == nil {
		return nil
	}

	var vendor *string
	if component.CPE != "" {
		cpe := ParseCPE(component.CPE)
		vendor = &cpe.Vendor
	}

	var description *string
	if component.Description != "" {
		description = &component.Description
	}
	size := int64(0)

	if component.Properties != nil {
		for _, prop := range *component.Properties {
			if prop.Name == "syft:metadata:installedSize" {
				s, err := strconv.Atoi(prop.Value)
				if err == nil {
					size = int64(s)
				}
			}
		}
	}

	return &Package{
		Sysinfo: s.NewSysinfo(TYPE_PACKAGE),
		Package: &PackageDetails{
			Architecture: &arch,
			Name:         &component.Name,
			Format:       pkg,
			Size:         &size,
			Version:      &component.Version,
			Vendor:       vendor,
			Description:  description,
		},
	}
}

func (s *SysCollector) SetPort(name string, port *PortInfo) {
	s.mxScan.Lock()
	defer s.mxScan.Unlock()
	s.localPorts[name] = port
}

func (s *SysCollector) RemovePort(name string) {
	s.mxScan.Lock()
	defer s.mxScan.Unlock()
	delete(s.localPorts, name)
}

var CpeNamePattern = regexp.MustCompile(`^[c][pP][eE]:(2\.3:|/)([AHOaho])?(.*)$`)

type CPE struct {
	Name            string  `xml:"name,attr" json:"name,omitempty"`
	CpeVersion      float32 `json:"cpe_version,omitempty"`
	Part            string  `json:"part,omitempty"`
	Vendor          string  `json:"vendor,omitempty"`
	Product         string  `json:"product,omitempty"`
	Version         string  `json:"version,omitempty"`
	Update          string  `json:"update,omitempty"`
	Edition         string  `json:"edition,omitempty"`
	Language        string  `json:"language,omitempty"`
	SoftwareEdition string  `json:"software_edition,omitempty"`
	TargetSoftware  string  `json:"target_software,omitempty"`
	TargetHardware  string  `json:"target_hardware,omitempty"`
	Other           string  `json:"other,omitempty"`
	Deprecated      bool    `xml:"deprecated,attr,omitempty" json:"deprecated,omitempty"`
	Title           string  `json:"title"`
}

func ParseCPE(name string) *CPE {
	cpe := &CPE{}
	// https://nvd.nist.gov/products/cpe
	if strings.HasPrefix(name, "cpe:2.3") {
		cpe.CpeVersion = 2.3
	} else {
		cpe.CpeVersion = 2.2
	}

	matches := CpeNamePattern.FindAllStringSubmatch(name, -1)
	var parts []string
	if len(matches) == 1 {
		match := matches[0]
		cpe.Part = match[2]
		parts = strings.Split(match[3], ":")

		if match[1] == "/" && len(parts) > 5 {
			moreParts := strings.Split(parts[5], "~")
			parts = parts[0:4]
			parts = append(parts, moreParts...)
		}

		for i, v := range parts {
			decoded, err := url.QueryUnescape(v)
			if err == nil {
				if decoded == "*" {
					parts[i] = ""
				} else {
					parts[i] = decoded
				}
			}
		}

		if len(parts) > 1 {
			cpe.Vendor = parts[1]
			if len(parts) > 2 {
				cpe.Product = parts[2]
				if len(parts) > 3 {
					cpe.Version = parts[3]
					if len(parts) > 4 {
						cpe.Update = parts[4]
						if len(parts) > 6 {
							cpe.Language = parts[6]
							if len(parts) > 7 {
								cpe.SoftwareEdition = parts[7]
								if len(parts) > 8 {
									cpe.TargetHardware = parts[8]
									if len(parts) > 9 {
										cpe.Other = parts[9]
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return cpe
}
