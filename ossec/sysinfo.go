package ossec

// See: https://github.com/wazuh/wazuh/blob/master/src/analysisd/decoders/syscollector.c

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	// https://www.socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var scanId uint64

func init() {
	scanId = uint64(time.Now().Unix())
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
	TYPE_PROCESS     = "process"
	TYPE_PROCESS_END = "process_end"
)

func (a *Client) PostSysinfo(input chan *QueuePosting) {
	scanId++
	input <- &QueuePosting{
		Location:    WM_SYS_LOCATION,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         a.NewOS(),
	}

	input <- &QueuePosting{
		Location:    WM_SYS_LOCATION,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         a.NewProcess(),
	}

	input <- &QueuePosting{
		Location:    WM_SYS_LOCATION,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         a.NewHardware(),
	}
}

type Sysinfo struct {
	// ScanTime  string     `json:"scan_time,omitempty"`
	Type     string `json:"type"`
	Checksum string `json:"checksum"`
	ID       string `json:"ID,omitempty"`
	ScanTime string `json:"scan_time,omitempty"`
}

func getScanTime() string {
	return time.Now().Format("2006/1/2 03:04:05")
}

func (a *Client) NewSysinfo(typ string, data interface{}) *Sysinfo {
	b, _ := json.Marshal(data)
	return &Sysinfo{
		Type:     typ,
		Checksum: fmt.Sprintf("%x", sha1.Sum(b)),
		ID:       fmt.Sprintf("%d", scanId),
		ScanTime: getScanTime(),
	}
}

type Hardware struct {
	// R"({"board_serial":"Intel Corporation","scan_time":"2020/12/28 21:49:50", "cpu_MHz":2904,"cpu_cores":2,"cpu_name":"Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz", "ram_free":2257872,"ram_total":4972208,"ram_usage":54})")));
	*Sysinfo
	BoardSerial string  `json:"board_serial,omitempty"`
	CPUName     string  `json:"cpu_name,omitempty"`
	CPUCores    int     `json:"cpu_cores,omitempty"`
	CPUMhz      float64 `json:"cpu_mhz,omitempty"`
	RamFree     uint64  `json:"ram_free,omitempty"`
	RamTotal    uint64  `json:"ram_total,omitempty"`
	RamUsage    float64 `json:"ram_usage,omitempty"`
}

func (a *Client) NewHardware() *Hardware {
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	hw := &Hardware{
		CPUName:  cpuStat[0].ModelName,
		CPUCores: runtime.NumCPU(),
		CPUMhz:   cpuStat[0].Mhz,
		RamFree:  vmStat.Free,
		RamTotal: vmStat.Total,
		RamUsage: vmStat.UsedPercent,
	}
	hw.Sysinfo = a.NewSysinfo(TYPE_OS, hw)
	return hw
}

type OS struct {
	//   R"({"architecture":"x86_64","scan_time":"2020/12/28 21:49:50", "hostname":"UBUNTU","os_build":"7601","os_major":"6","os_minor":"1","os_name":"Microsoft Windows 7","os_release":"sp1","os_version":"6.1.7601"})")));
	*Sysinfo
	ScanTime     string `json:"scan_time,omitempty"`
	OSName       string `json:"os_name,omitempty"`
	OSMajor      string `json:"os_major,omitempty"`
	OSMinor      string `json:"os_minor,omitempty"`
	OSBuild      string `json:"os_build,omitempty"`
	OSVersion    string `json:"os_version,omitempty"`
	Hostname     string `json:"hostname,omitempty"`
	OSRelease    string `json:"os_release,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

func (a *Client) NewOS() *OS {
	os := &OS{
		OSName:       a.osInfo.Name,
		OSVersion:    a.osInfo.Version,
		Hostname:     a.un.Hostname,
		OSRelease:    a.osInfo.Release,
		Architecture: runtime.GOARCH,
	}
	os.Sysinfo = a.NewSysinfo(TYPE_OS, os)
	return os
}

// R"([{"architecture":"amd64","scan_time":"2020/12/28 21:49:50", "group":"x11","name":"xserver-xorg","priority":"optional","size":"411","source":"xorg","version":"1:7.7+19ubuntu14","os_patch":""},{"hotfix":"KB4586786"}])")));

type Process struct {
	// R"([{"egroup":"root","euser":"root","fgroup":"root","name":"kworker/u256:2-","scan_time":"2020/12/28 21:49:50", "nice":0,"nlwp":1,"pgrp":0,"pid":431625,"ppid":2,"priority":20,"processor":1,"resident":0,"rgroup":"root","ruser":"root","session":0,"sgroup":"root","share":0,"size":0,"start_time":9302261,"state":"I","stime":3,"suser":"root","tgid":431625,"tty":0,"utime":0,"vm_size":0}])")));

	*Sysinfo
	Name      string `json:"name,omitempty"`
	EGroup    string `json:"egroup,omitempty"`
	EUser     string `json:"euser,omitempty"`
	FGroup    string `json:"fgroup,omitempty"`
	RGroup    string `json:"rgroup,omitempty"`
	RUser     string `json:"ruser,omitempty"`
	SGroup    string `json:"sgroup,omitempty"`
	SUser     string `json:"suser,omitempty"`
	State     string `json:"state,omitempty"`
	Nice      int    `json:"nice,omitempty"`
	NLWP      int    `json:"nlwp,omitempty"`
	PGrp      int    `json:"pgrp,omitempty"`
	PID       int    `json:"pid,omitempty"`
	PPID      int    `json:"ppid,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	Processor int    `json:"processor,omitempty"`
	Resident  int    `json:"resident,omitempty"`
	Session   int    `json:"session,omitempty"`
	Share     int    `json:"share,omitempty"`
	Size      int    `json:"size,omitempty"`
	STime     int    `json:"stime,omitempty"`
	TGID      int    `json:"tgid,omitempty"`
	TTY       int    `json:"tty,omitempty"`
	UTime     int    `json:"utime,omitempty"`
	VMSize    int    `json:"vm_size,omitempty"`
}

func (a *Client) NewProcess() *Process {
	p := &Process{
		Name: filepath.Base(os.Args[0]),
		PID:  os.Getpid(),
		PPID: os.Getppid(),
	}

	p.Sysinfo = a.NewSysinfo(TYPE_OS, p)
	return p
}

// R"([{"architecture":"amd64","scan_time":"2020/12/28 21:49:50", "group":"x11","name":"xserver-xorg","priority":"optional","size":"411","source":"xorg","version":"1:7.7+19ubuntu14","os_patch":""},{"hotfix":"KB4586786"}])")));
type Package struct {
	*Sysinfo
	Architecture string `json:"architecture,omitempty"`
	Group        string `json:"group,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     string `json:"priority,omitempty"`
	Size         string `json:"size,omitempty"`
	Source       string `json:"source,omitempty"`
	Version      string `json:"version,omitempty"`
	OSPatch      string `json:"os_patch,omitempty"`
	Hotfix       string `json:"hotfix,omitempty"`
}

type Network struct {
	// R"({"iface":[{"address":"127.0.0.1","scan_time":"2020/12/28 21:49:50", "mac":"d4:5d:64:51:07:5d", "gateway":"192.168.0.1|600","broadcast":"127.255.255.255", "name":"ens1", "mtu":1500, "name":"enp4s0", "adapter":" ", "type":"ethernet", "state":"up", "dhcp":"disabled","iface":"Loopback Pseudo-Interface 1","metric":"75","netmask":"255.0.0.0","proto":"IPv4","rx_bytes":0,"rx_dropped":0,"rx_errors":0,"rx_packets":0,"tx_bytes":0,"tx_dropped":0,"tx_errors":0,"tx_packets":0, "IPv4":[{"address":"192.168.153.1","broadcast":"192.168.153.255","dhcp":"unknown","metric":" ","netmask":"255.255.255.0"}], "IPv6":[{"address":"fe80::250:56ff:fec0:8","dhcp":"unknown","metric":" ","netmask":"ffff:ffff:ffff:ffff::"}]}]})")));

	*Sysinfo
}

type Port struct {
	// R"({"ports":[{"inode":0,"local_ip":"127.0.0.1","scan_time":"2020/12/28 21:49:50", "local_port":631,"pid":0,"process_name":"System Idle Process","protocol":"tcp","remote_ip":"0.0.0.0","remote_port":0,"rx_queue":0,"state":"listening","tx_queue":0}]})")));
	*Sysinfo
}
