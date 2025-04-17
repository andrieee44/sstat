package sstat_test

import (
	"fmt"

	"github.com/andrieee44/sstat"
)

func ExamplePowerSupply() {
	var (
		powerSupplyInfo *sstat.PowerSupplyInfo
		err             error
	)

	powerSupplyInfo, err = sstat.PowerSupply("ADP0")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", powerSupplyInfo)
}

func ExamplePowerSupplies() {
	var (
		powerSupplyInfos []*sstat.PowerSupplyInfo
		idx              int
		err              error
	)

	powerSupplyInfos, err = sstat.PowerSupplies("ADP*")
	if err != nil {
		panic(err)
	}

	for idx = range powerSupplyInfos {
		fmt.Printf("%+v\n", powerSupplyInfos[idx])
	}
}

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
