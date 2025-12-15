// Copyright IBM Corp. 2016, 2025

package libcontainer

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestErrorDetail(t *testing.T) {
	err := newGenericError(fmt.Errorf("test error"), SystemError)
	if derr := err.Detail(ioutil.Discard); derr != nil {
		t.Fatal(derr)
	}
}
