// Copyright IBM Corp. 2016, 2025

package configs

type HugepageLimit struct {
	// which type of hugepage to limit.
	Pagesize string `json:"page_size"`

	// usage limit for hugepage.
	Limit uint64 `json:"limit"`
}
