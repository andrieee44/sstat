package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

// Get the name of ADP0.
func ExamplePowerSupply() {
	var (
		powerSupplyInfo *sstat.PowerSupplyInfo
		err             error
	)

	powerSupplyInfo, err = sstat.PowerSupply("ADP0")
	if err != nil {
		panic(err)
	}

	fmt.Println(powerSupplyInfo.Name())
}

// Get the type of all power supplies available.
func ExamplePowerSupplies() {
	var (
		powerSupplyInfos []*sstat.PowerSupplyInfo
		idx              int
		err              error
	)

	powerSupplyInfos, err = sstat.PowerSupplies("*")
	if err != nil {
		panic(err)
	}

	for idx = range powerSupplyInfos {
		fmt.Println(powerSupplyInfos[idx].Type())
	}
}

// Get the power supply manufacturer of ADP0.
func ExamplePowerSupplyInfo_UeventKey() {
	var (
		powerSupplyInfo *sstat.PowerSupplyInfo
		value           string
		ok              bool
		err             error
	)

	powerSupplyInfo, err = sstat.PowerSupply("ADP0")
	if err != nil {
		panic(err)
	}

	value, ok = powerSupplyInfo.UeventKey("POWER_SUPPLY_MANUFACTURER")
	if !ok {
		panic(fmt.Errorf("%s: key not found in uevent file", "POWER_SUPPLY_MANUFACTURER"))
	}

	fmt.Println(value)
}
