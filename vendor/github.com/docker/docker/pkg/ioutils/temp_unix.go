// Copyright IBM Corp. 2016, 2025

// +build !windows

package ioutils

import "io/ioutil"

// TempDir on Unix systems is equivalent to ioutil.TempDir.
func TempDir(dir, prefix string) (string, error) {
	return ioutil.TempDir(dir, prefix)
}
