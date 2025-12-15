// Copyright IBM Corp. 2016, 2025

// +build !windows

package libnetwork

import "github.com/docker/libnetwork/ipamapi"

// Stub implementations for DNS related functions

func (n *network) startResolver() {
}

func defaultIpamForNetworkType(networkType string) string {
	return ipamapi.DefaultIPAM
}
