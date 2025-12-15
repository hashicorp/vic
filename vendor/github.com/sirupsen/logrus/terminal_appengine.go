// Copyright IBM Corp. 2016, 2025

// +build appengine

package logrus

// IsTerminal returns true if stderr's file descriptor is a terminal.
func IsTerminal() bool {
	return true
}
