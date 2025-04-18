package sstat_test

import (
	"fmt"
	"path/filepath"

	"github.com/andrieee44/sstat"
)

// Get the current battery operating status.
func ExamplePathReadStr() {
	var (
		status string
		err    error
	)

	status, err = sstat.PathReadStr(filepath.Join(sstat.PowerSupplyPath, "BAT0", "status"))
	if err != nil {
		panic(err)
	}

	fmt.Println("BAT0 status:", status)
}

// Get the current remaining battery percentage.
func ExamplePathReadInt() {
	var (
		capacity int
		err      error
	)

	capacity, err = sstat.PathReadInt(filepath.Join(sstat.PowerSupplyPath, "BAT0", "capacity"))
	if err != nil {
		panic(err)
	}

	fmt.Println("BAT0 percentage:", capacity)
}
