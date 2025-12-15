// Copyright IBM Corp. 2016, 2025

// +build !linux

package archive

func getWhiteoutConverter(format WhiteoutFormat) tarWhiteoutConverter {
	return nil
}
