// Copyright IBM Corp. 2016, 2025

// +build !exclude_graphdriver_zfs,linux !exclude_graphdriver_zfs,freebsd, solaris

package register

import (
	// register the zfs driver
	_ "github.com/docker/docker/daemon/graphdriver/zfs"
)
