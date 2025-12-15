// Copyright IBM Corp. 2016, 2025

// +build !cgo,!plan9 windows android

package sftp

import (
	"os"
)

func fileStatFromInfoOs(fi os.FileInfo, flags *uint32, fileStat *FileStat) {
	// todo
}
