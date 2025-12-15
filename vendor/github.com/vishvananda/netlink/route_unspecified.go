// Copyright IBM Corp. 2016, 2025

// +build !linux

package netlink

func (r *Route) ListFlags() []string {
	return []string{}
}
