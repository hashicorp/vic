// Copyright IBM Corp. 2016, 2025

package dhcp4client

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
