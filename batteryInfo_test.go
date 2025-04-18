package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

// Get battery information for BAT0.
func ExampleBattery() {
	var (
		batteryInfo *sstat.BatteryInfo
		err         error
	)

	batteryInfo, err = sstat.Battery("BAT0")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", batteryInfo)
}

// Get all battery information.
func ExampleBatteries() {
	var (
		batteryInfos []*sstat.BatteryInfo
		idx          int
		err          error
	)

	batteryInfos, err = sstat.Batteries()
	if err != nil {
		panic(err)
	}

	for idx = range batteryInfos {
		fmt.Printf("%+v\n", batteryInfos[idx])
	}
}
