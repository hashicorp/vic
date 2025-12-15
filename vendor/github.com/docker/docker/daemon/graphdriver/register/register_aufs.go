// Copyright IBM Corp. 2016, 2025

// +build !exclude_graphdriver_aufs,linux

package register

import (
	// register the aufs graphdriver
	_ "github.com/docker/docker/daemon/graphdriver/aufs"
)
