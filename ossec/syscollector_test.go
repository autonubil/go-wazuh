package ossec

import (
	"testing"

	"github.com/autonubil/go-wazuh/sysinfo"
)

// Successfully sends OS info to input channel
func TestSendPackages_SuccessfullySendsOSInfo(t *testing.T) {
	// PostSysinfo emits OS, Hardware, network/port/package entries and process
	// info in one synchronous call; buffer generously so it never blocks on an
	// unread channel.
	input := make(chan *QueuePosting, 64)

	a, err := NewAgent("127.0.0.1", "666", "zbook-cze", "c588215eddc469001c41f4aa5cc306ea20fc6815e5cfe3cc666c5ccb7bc505e0",
		WithAgentAllowedIPs("any"), WithEncryptionMethod(EncryptionMethodBlowFish),
		WithOSSettings(&sysinfo.OS{Name: "TestOS", Version: "1.0"}))
	if err != nil {
		t.Error(err)
		return
	}

	a.Scanner.PostSysinfo(input)

	// Assert that the OS info is sent to the input channel
	select {
	case posting := <-input:
		osInfo, ok := posting.Raw.(*OS)
		if !ok {
			t.Errorf("Expected *OS, got %T", posting.Raw)
		}
		if osInfo.Inventory.OSName == nil || *osInfo.Inventory.OSName != "TestOS" {
			t.Errorf("Expected OS name to be 'TestOS', got %v", osInfo.Inventory.OSName)
		}
		if osInfo.Inventory.OSVersion == nil || *osInfo.Inventory.OSVersion != "1.0" {
			t.Errorf("Expected OS version to be '1.0', got %v", osInfo.Inventory.OSVersion)
		}
		// Hostname comes from GetMachineNameSafe(), not from WithOSSettings — assert
		// it's populated rather than pinning it to a value nothing sets.
		if osInfo.Inventory.Hostname == nil || *osInfo.Inventory.Hostname == "" {
			t.Errorf("Expected a non-empty hostname, got %v", osInfo.Inventory.Hostname)
		}
	default:
		t.Error("Expected an OS info posting, but no posting was received")
	}
}

// Unable to get current working directory
func TestSendPackages_UnableToGetCurrentWorkingDirectory(t *testing.T) {
	// Create a mock input channel
	input := make(chan *QueuePosting)

	// Create a mock SysLinuxCollector instance
	sysCollector := &SysLinuxCollector{
		agent: &Client{},
	}

	// Call the sendPackages function
	sysCollector.sendPackages(input)

	// Assert that no SBOM postings are sent to the input channel
	select {
	case posting := <-input:
		pkg, ok := posting.Raw.(*Package)
		if !ok {
			t.Errorf("Expected *Package, got %T", posting.Raw)
		}
		if pkg.Sysinfo.Type != TYPE_PACKAGE_END {
			t.Errorf("Expected package type to be 'program_end', got %s", pkg.Sysinfo.Type)
		}
		t.Error("Expected no SBOM postings, but a posting was received")
	default:
		// No posting received, which is expected
	}
}
