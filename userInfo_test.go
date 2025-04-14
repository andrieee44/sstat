package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

func ExampleCurrentUser() {
	var (
		userInfo *sstat.UserInfo
		err      error
	)

	userInfo, err = sstat.CurrentUser()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", userInfo)
}

func ExampleLookupUser() {
	var (
		userInfo *sstat.UserInfo
		err      error
	)

	userInfo, err = sstat.LookupUser("root")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", userInfo)
}

func ExampleLookupUserId() {
	var (
		userInfo *sstat.UserInfo
		err      error
	)

	userInfo, err = sstat.LookupUserId("0")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", userInfo)
}
