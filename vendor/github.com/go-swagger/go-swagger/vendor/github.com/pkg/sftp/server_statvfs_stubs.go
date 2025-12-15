// Copyright IBM Corp. 2016, 2025

// +build !darwin,!linux gccgo

package sftp

import (
	"syscall"
)

func (p sshFxpExtendedPacketStatVFS) respond(svr *Server) error {
	return syscall.ENOTSUP
}
