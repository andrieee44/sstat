package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

// Print the type of ADP0.
func ExamplePowerSupply() {
	var (
		powerSupplyInfo *sstat.PowerSupplyInfo
		err             error
	)

	powerSupplyInfo, err = sstat.PowerSupply("ADP0")
	if err != nil {
		panic(err)
	}

	fmt.Println(powerSupplyInfo.Type())
}

// Print the names of all power supplies available.
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
		fmt.Println(powerSupplyInfos[idx].Name())
	}
}

// Print the power supply manufacturer of ADP0.
func ExamplePowerSupplyInfo_Key() {
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

	value, ok = powerSupplyInfo.Key("POWER_SUPPLY_MANUFACTURER")
	if !ok {
		panic(fmt.Errorf("%s: key not found in uevent file", "POWER_SUPPLY_MANUFACTURER"))
	}

	fmt.Println(value)
}
