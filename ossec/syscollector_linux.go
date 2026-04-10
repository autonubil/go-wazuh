//go:build linux && !android

package ossec

// See: https://documentation.wazuh.com/current/user-manual/capabilities/SysLinuxCollector.html
// See: https://github.com/wazuh/wazuh/blob/master/src/analysisd/decoders/SysLinuxCollector.c
// See: https://github.com/wazuh/wazuh/blob/master/framework/wazuh/core/SysLinuxCollector.py

import (
	"encoding/gob"
	"net"
	"os"
	"os/user"
	"path/filepath"
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

func init() {
	gob.Register(SysLinuxCollector{})
}

type SysLinuxCollector struct {
	agent      *Client
	scanId     int64
	localPorts map[string]*PortInfo
	scanTime   string
	mxScan     sync.Mutex
}

func NewPlatformScanner(client *Client) SysCollector {
	return &SysLinuxCollector{
		scanId:     int64(time.Now().Unix()),
		agent:      client,
		localPorts: make(map[string]*PortInfo),
	}
}

func (s *SysLinuxCollector) PostSysinfo(input chan *QueuePosting) {
	s.mxScan.Lock()
	defer s.mxScan.Unlock()
	s.scanId++
	s.scanTime = getScanTime()

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         s.NewOS(s.scanId, s.scanTime),
	}

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         s.NewHardware(s.scanId, s.scanTime),
	}

	s.sendNetworks(input)
	s.sendPorts(input)
	s.sendPackages(input)

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         s.NewProcess(TYPE_PROCESS, s.scanId, s.scanTime),
	}

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         &Process{Sysinfo: NewSysinfo(TYPE_PROCESS_END, s.scanId, s.scanTime)},
	}

}

func (s *SysLinuxCollector) sendPackages(input chan *QueuePosting) {
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
				s.sendSBOM(input, path, s.scanId, s.scanTime)
			}
			return nil
		})
	}
}

func (s *SysLinuxCollector) sendSBOM(input chan *QueuePosting, path string, scanId int64, scanTime string) {
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
		pkg := NewPackageFromComponent(component, scanId, scanTime)
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
		Raw:         &Package{Sysinfo: NewSysinfo(TYPE_PACKAGE_END, scanId, scanTime)},
	}

}

func (s *SysLinuxCollector) sendPorts(input chan *QueuePosting) {
	for _, pi := range s.localPorts {
		input <- &QueuePosting{
			Location:    SYSCOLLECTOR_MOD,
			TargetQueue: SYSCOLLECTOR_MQ,
			Timestamp:   time.Now(),
			Raw:         &Port{Sysinfo: NewSysinfo(TYPE_PORT, s.scanId, s.scanTime), PortInfo: s.NewPortInfo(pi)},
		}
	}

	input <- &QueuePosting{
		Location:    SYSCOLLECTOR_MOD,
		TargetQueue: SYSCOLLECTOR_MQ,
		Timestamp:   time.Now(),
		Raw:         &Port{Sysinfo: NewSysinfo(TYPE_PORT_END, s.scanId, s.scanTime)},
	}
}

func (s *SysLinuxCollector) sendNetworks(input chan *QueuePosting) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, iface := range ifaces {
		net, isLocal := s.NewNetwork(iface, s.scanId, s.scanTime)
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
		Raw:         &Network{Sysinfo: NewSysinfo(TYPE_NETWORK_END, s.scanId, s.scanTime)},
	}

}

func (s *SysLinuxCollector) NewHardware(scanId int64, scanTime string) *Hardware {
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	ramFree := vmStat.Free / 1024
	ramTotal := vmStat.Total / 1024
	ramUsage := vmStat.UsedPercent / 1024

	numCpu := runtime.NumCPU()
	hw := &Hardware{
		Sysinfo: NewSysinfo(TYPE_HARDWARE, scanId, scanTime),
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

func (s *SysLinuxCollector) NewK8sNodeOS(agent *Client, scanId int64, scanTime string) *OS {
	result := NewOS(agent, scanId, scanTime)

	_, err := os.Stat("/var/run/secrets/kubernetes.io")
	if err == nil {
		// Get the k8s node name from the environment variable
		nodeName := os.Getenv("NODE_NAME")
		if nodeName != "" {
			result.Inventory.Hostname = &nodeName
		}
		nodeName = os.Getenv("K8S_NODE_NAME")
		if nodeName != "" {
			result.Inventory.Hostname = &nodeName
		}
	}
	return result
}

func (s *SysLinuxCollector) NewOS(scanId int64, scanTime string) *OS {
	arch := runtime.GOARCH
	os := &OS{
		Sysinfo: NewSysinfo(TYPE_OS, scanId, scanTime),
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

func (s *SysLinuxCollector) NewProcess(typ string, scanId int64, scanTime string) *Process {
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
		Sysinfo: NewSysinfo(typ, scanId, scanTime),
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

func (s *SysLinuxCollector) NewNetwork(intf net.Interface, scanId int64, scanTime string) (*Network, bool) {
	isLocal := false
	eth := "ethernet"
	state := "down"
	if intf.Flags|net.FlagUp == intf.Flags {
		state = "up"
	}
	network := &Network{
		Sysinfo: NewSysinfo(TYPE_NETWORK, scanId, scanTime),
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

func (s *SysLinuxCollector) NewPort(pi *PortInfo, scanId int64, scanTime string) *Port {
	cmd := os.Args[0]
	pid := os.Getpid()

	pi.Pid = &pid
	pi.Process = &cmd

	port := &Port{
		Sysinfo:  NewSysinfo(TYPE_PORT, scanId, scanTime),
		PortInfo: pi,
	}
	return port
}

func (s *SysLinuxCollector) NewPortInfo(pi *PortInfo) *PortInfo {
	cmd := os.Args[0]
	pid := os.Getpid()

	pi.Pid = &pid
	pi.Process = &cmd

	return pi
}

func (s *SysLinuxCollector) SetPort(name string, port *PortInfo) {
	s.mxScan.Lock()
	defer s.mxScan.Unlock()
	s.localPorts[name] = port
}

func (s *SysLinuxCollector) RemovePort(name string) {
	s.mxScan.Lock()
	defer s.mxScan.Unlock()
	delete(s.localPorts, name)
}
