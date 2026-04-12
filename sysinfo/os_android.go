//go:build android

package sysinfo

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// GetOSInfo returns android specific information
func GetOSInfo() *OS {
	osInfo := &OS{}

	// Detect Architecture using Go runtime for better reliability on ARM/ARM64
	osInfo.Architecture = runtime.GOARCH

	// Check if we are actually on Android
	if _, err := os.Stat("/system/build.prop"); err == nil {
		osInfo.getAndroidInfo()
	} else {
		// Fallback or placeholder for non-android environments
		osInfo.Name = "Unknown (Non-Android)"
	}

	return osInfo
}

// readProperty fetches values from Android's system properties
func readProperty(prop string) string {
	out, err := exec.Command("/system/bin/getprop", prop).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func (osInfo *OS) getAndroidInfo() {
	osInfo.Name = "Android"
	osInfo.Vendor = readProperty("ro.product.manufacturer")
	// Android version (e.g., 13)
	osInfo.Version = readProperty("ro.build.version.release")
	// SDK level (e.g., 33)
	osInfo.Release = readProperty("ro.build.version.sdk")
	// Build ID
	osInfo.Build = readProperty("ro.build.display.id")
	// Internal codename
	osInfo.Codename = readProperty("ro.build.version.codename")
}
