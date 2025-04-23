package sstat

import (
	"testing"
)

func TestPowerSupply(t *testing.T) {
	var err error

	if !checkBat(t) {
		return
	}

	_, err = PowerSupply("BAT0")
	tErrorIf(t, err)
}

func TestPowerSupplies(t *testing.T) {
	var err error

	if !checkBat(t) {
		return
	}

	_, err = PowerSupplies("BAT*")
	tErrorIf(t, err)
}
