// Copyright IBM Corp. 2016, 2025

// +build go1.4

package handlers

import (
	"net/http"
)

func basicAuth(r *http.Request) (username, password string, ok bool) {
	return r.BasicAuth()
}
