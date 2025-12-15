// Copyright IBM Corp. 2016, 2025

// +build !cgo,!plan9 windows android

package sftp

import (
	"os"
	"path"
)

func runLs(dirname string, dirent os.FileInfo) string {
	return path.Join(dirname, dirent.Name())
}
