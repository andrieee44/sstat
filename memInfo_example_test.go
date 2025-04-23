package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

// Print the current memory percentage.
func ExampleNewMemInfo() {
	var (
		memInfo *sstat.MemInfo
		err     error
	)

	memInfo, err = sstat.NewMemInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(memInfo)
}
