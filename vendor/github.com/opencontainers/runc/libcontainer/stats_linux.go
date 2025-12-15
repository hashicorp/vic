// Copyright IBM Corp. 2016, 2025

package libcontainer

import "github.com/opencontainers/runc/libcontainer/cgroups"

type Stats struct {
	Interfaces  []*NetworkInterface
	CgroupStats *cgroups.Stats
}
