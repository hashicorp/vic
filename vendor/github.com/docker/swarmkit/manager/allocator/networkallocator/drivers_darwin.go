// Copyright IBM Corp. 2016, 2025

package networkallocator

import (
	"github.com/docker/libnetwork/drivers/overlay/ovmanager"
	"github.com/docker/libnetwork/drivers/remote"
)

func getInitializers() []initializer {
	return []initializer{
		{remote.Init, "remote"},
		{ovmanager.Init, "overlay"},
	}
}
