//go:build android

package ossec

// See: https://documentation.wazuh.com/current/user-manual/capabilities/SysAndroidCollector.html
// See: https://github.com/wazuh/wazuh/blob/master/src/analysisd/decoders/SysAndroidCollector.c
// See: https://github.com/wazuh/wazuh/blob/master/framework/wazuh/core/SysAndroidCollector.py

import (
	"encoding/gob"
	"sync"
	"time"
	// https://www.socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
)

func init() {
	gob.Register(SysAndroidCollector{})
}

type SysAndroidCollector struct {
	agent      *Client
	scanId     int64
	localPorts map[string]*PortInfo
	scanTime   string
	mxScan     sync.Mutex
}

func NewPlatformScanner(client *Client) SysCollector {
	return &SysAndroidCollector{
		scanId:     int64(time.Now().Unix()),
		agent:      client,
		localPorts: make(map[string]*PortInfo),
	}
}

func (s *SysAndroidCollector) PostSysinfo(input chan *QueuePosting) {
	s.mxScan.Lock()
	defer s.mxScan.Unlock()
	s.scanId++
	s.scanTime = getScanTime()
	/*
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
	*/
}
