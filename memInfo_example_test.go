package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

// Print the current memory percentage.
func ExampleNewMemInfo() {
	var (
		memInfo                            *sstat.MemInfo
		memTotal, memFree, buffers, cached int
		err                                error
	)

	memInfo, err = sstat.NewMemInfo()
	if err != nil {
		panic(err)
	}

	memInfo.Populate(map[string]*int{
		"MemTotal": &memTotal,
		"MemFree":  &memFree,
		"Buffers":  &buffers,
		"Cached":   &cached,
	})

	fmt.Printf("Used memory: %g%%\n", float64(memTotal-memFree-buffers-cached)/float64(memTotal)*100)
}
