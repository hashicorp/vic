// Copyright IBM Corp. 2016, 2025

// +build !windows

package cobra

var preExecHookFn func(*Command) = nil
