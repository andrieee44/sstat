package sstat

import "testing"

func TestCurrentUser(t *testing.T) {
	var err error

	_, err = CurrentUser()
	tErrorIf(t, err)
}

func TestLookupUser(t *testing.T) {
	var err error

	_, err = LookupUser("root")
	tErrorIf(t, err)
}

func TestLookupUserId(t *testing.T) {
	var err error

	_, err = LookupUserId("0")
	tErrorIf(t, err)
}
