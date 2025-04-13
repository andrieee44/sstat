package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

func ExampleGetCurrentUser() {
	var (
		currentUser *sstat.UserInfo
		err         error
	)

	currentUser, err = sstat.GetCurrentUser()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", currentUser)
}
