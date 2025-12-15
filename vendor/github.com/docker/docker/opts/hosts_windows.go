// Copyright IBM Corp. 2016, 2025

// +build windows

package opts

// DefaultHost constant defines the default host string used by docker on Windows
var DefaultHost = "npipe://" + DefaultNamedPipe
