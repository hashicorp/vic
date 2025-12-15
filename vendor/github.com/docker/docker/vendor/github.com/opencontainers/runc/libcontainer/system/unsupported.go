// Copyright IBM Corp. 2016, 2025

// +build !linux

package system

// RunningInUserNS is a stub for non-Linux systems
// Always returns false
func RunningInUserNS() bool {
	return false
}
