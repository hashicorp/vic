// Copyright IBM Corp. 2016, 2025

// +build windows

package sysinfo

// New returns an empty SysInfo for windows for now.
func New(quiet bool) *SysInfo {
	sysInfo := &SysInfo{}
	return sysInfo
}
