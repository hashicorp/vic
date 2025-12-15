// Copyright IBM Corp. 2016, 2025

// +build cgo

package metrics

import "runtime"

func numCgoCall() int64 {
	return runtime.NumCgoCall()
}
