// Copyright IBM Corp. 2016, 2025

// +build cgo,linux cgo,freebsd

package system

/*
#include <unistd.h>
*/
import "C"

func GetClockTicks() int {
	return int(C.sysconf(C._SC_CLK_TCK))
}
