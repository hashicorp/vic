// Copyright IBM Corp. 2016, 2025

package archive

import (
	"path/filepath"
)

func normalizePath(path string) string {
	return filepath.FromSlash(path)
}
