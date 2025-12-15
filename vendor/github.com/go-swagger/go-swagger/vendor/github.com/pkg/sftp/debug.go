// Copyright IBM Corp. 2016, 2025

// +build debug

package sftp

import "log"

func debug(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}
