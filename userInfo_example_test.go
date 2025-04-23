package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

// Get the username of the current user.
func ExampleCurrentUser() {
	var (
		userInfo *sstat.UserInfo
		err      error
	)

	userInfo, err = sstat.CurrentUser()
	if err != nil {
		panic(err)
	}

	fmt.Println(userInfo.Username())
}

// Get the userid of root.
func ExampleLookupUser() {
	var (
		userInfo *sstat.UserInfo
		err      error
	)

	userInfo, err = sstat.LookupUser("root")
	if err != nil {
		panic(err)
	}

	fmt.Println(userInfo.Uid())
}

// Get the group name of userid 0.
func ExampleLookupUserId() {
	var (
		userInfo *sstat.UserInfo
		err      error
	)

	userInfo, err = sstat.LookupUserId("0")
	if err != nil {
		panic(err)
	}

	fmt.Println(userInfo.Group())
}
