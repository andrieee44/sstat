package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

// Print the battery capacity percent of BAT0.
func ExampleBattery() {
	var (
		batteryInfo *sstat.BatteryInfo
		err         error
	)

	batteryInfo, err = sstat.Battery("BAT0")
	if err != nil {
		panic(err)
	}

	fmt.Println(batteryInfo.Capacity())
}

// Print the names of all batteries.
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
		fmt.Println(batteryInfos[idx].Name())
	}
}
