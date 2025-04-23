package sstat

import "testing"

func TestMemInfo(t *testing.T) {
	var err error

	_, err = NewMemInfo()
	tErrorIf(t, err)
}
