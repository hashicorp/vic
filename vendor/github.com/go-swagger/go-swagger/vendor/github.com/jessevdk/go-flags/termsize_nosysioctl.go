// Copyright IBM Corp. 2016, 2025

// +build windows plan9 solaris

package flags

func getTerminalColumns() int {
	return 80
}
