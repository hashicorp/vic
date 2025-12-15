// Copyright IBM Corp. 2016, 2025

package fileutils

// GetTotalUsedFds Returns the number of used File Descriptors. Not supported
// on Windows.
func GetTotalUsedFds() int {
	return -1
}
