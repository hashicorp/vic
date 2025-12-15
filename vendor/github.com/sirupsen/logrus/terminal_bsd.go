// Copyright IBM Corp. 2016, 2025

// +build darwin freebsd openbsd netbsd dragonfly
// +build !appengine

package logrus

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA

type Termios syscall.Termios
