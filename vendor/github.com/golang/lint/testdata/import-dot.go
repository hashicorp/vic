// Copyright IBM Corp. 2016, 2025

// Test that dot imports are flagged.

// Package pkg ...
package pkg

import . "fmt" // MATCH /dot import/

var _ Stringer // from "fmt"
