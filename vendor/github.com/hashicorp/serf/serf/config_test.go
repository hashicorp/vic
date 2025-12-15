// Copyright IBM Corp. 2016, 2025

package serf

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	c := DefaultConfig()
	if c.ProtocolVersion != ProtocolVersionMax {
		t.Fatalf("bad: %#v", c)
	}
}
