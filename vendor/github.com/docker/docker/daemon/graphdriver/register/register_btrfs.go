// Copyright IBM Corp. 2016, 2025

// +build !exclude_graphdriver_btrfs,linux

package register

import (
	// register the btrfs graphdriver
	_ "github.com/docker/docker/daemon/graphdriver/btrfs"
)
