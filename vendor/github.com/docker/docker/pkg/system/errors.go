// Copyright IBM Corp. 2016, 2025

package system

import (
	"errors"
)

var (
	// ErrNotSupportedPlatform means the platform is not supported.
	ErrNotSupportedPlatform = errors.New("platform and architecture is not supported")
)
