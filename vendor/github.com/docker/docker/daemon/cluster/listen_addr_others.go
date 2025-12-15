// Copyright IBM Corp. 2016, 2025

// +build !linux,!solaris

package cluster

import "net"

func (c *Cluster) resolveSystemAddr() (net.IP, error) {
	return c.resolveSystemAddrViaSubnetCheck()
}
