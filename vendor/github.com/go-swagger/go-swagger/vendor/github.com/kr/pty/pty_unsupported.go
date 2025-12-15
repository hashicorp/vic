// Copyright IBM Corp. 2016, 2025

// +build !linux,!darwin,!freebsd,!dragonfly

package pty

import (
	"os"
)

func open() (pty, tty *os.File, err error) {
	return nil, nil, ErrUnsupported
}
