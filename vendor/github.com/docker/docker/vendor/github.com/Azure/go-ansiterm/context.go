// Copyright IBM Corp. 2016, 2025

package ansiterm

type ansiContext struct {
	currentChar byte
	paramBuffer []byte
	interBuffer []byte
}
