// Copyright IBM Corp. 2016, 2025

// +build !linux,!darwin,!freebsd,!windows,!solaris

package signal

import (
	"syscall"
)

// SignalMap is an empty map of signals for unsupported platform.
var SignalMap = map[string]syscall.Signal{}
