// Copyright IBM Corp. 2016, 2025

// +build !linux,!freebsd

package system

import "syscall"

// LUtimesNano is only supported on linux and freebsd.
func LUtimesNano(path string, ts []syscall.Timespec) error {
	return ErrNotSupportedPlatform
}
