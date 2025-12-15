// Copyright IBM Corp. 2016, 2025

package container

import (
	"golang.org/x/sys/unix"
)

func detachMounted(path string) error {
	return unix.Unmount(path, unix.MNT_DETACH)
}
