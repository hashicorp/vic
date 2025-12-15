// Copyright IBM Corp. 2016, 2025

package dbus

func (t *unixTransport) SendNullByte() error {
	_, err := t.Write([]byte{0})
	return err
}
