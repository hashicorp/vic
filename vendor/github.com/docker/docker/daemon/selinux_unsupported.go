// Copyright IBM Corp. 2016, 2025

// +build !linux

package daemon

func selinuxSetDisabled() {
}

func selinuxFreeLxcContexts(label string) {
}

func selinuxEnabled() bool {
	return false
}
