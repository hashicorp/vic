// Copyright IBM Corp. 2016, 2025

// +build !linux,!solaris,!windows

package sysinfo

// New returns an empty SysInfo for non linux nor solaris for now.
func New(quiet bool) *SysInfo {
	sysInfo := &SysInfo{}
	return sysInfo
}
