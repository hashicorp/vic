// Copyright IBM Corp. 2016, 2025

// +build experimental

package libnetwork

import "github.com/docker/libnetwork/drivers/ipvlan"

func additionalDrivers() []initializer {
	return []initializer{
		{ipvlan.Init, "ipvlan"},
	}
}
