// Copyright IBM Corp. 2016, 2025

package panicwrap

import "fmt"

func monitor(c *WrapConfig) (int, error) {
	return -1, fmt.Errorf("Monitor is not supported on windows")
}
