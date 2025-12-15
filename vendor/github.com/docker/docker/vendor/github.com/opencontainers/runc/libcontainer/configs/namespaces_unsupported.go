// Copyright IBM Corp. 2016, 2025

// +build !linux,!freebsd

package configs

// Namespace defines configuration for each namespace.  It specifies an
// alternate path that is able to be joined via setns.
type Namespace struct {
}
