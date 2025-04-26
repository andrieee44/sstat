package sstat

import (
	"path/filepath"
	"testing"
)

func TestPowerSupply(t *testing.T) {
	var err error

	if !checkPath(t, filepath.Join(PowerSupplyPath, "BAT0")) {
		return
	}

	_, err = PowerSupply("BAT0")
	tErrorIf(t, err)
}

func TestPowerSupplies(t *testing.T) {
	var err error

	if !checkPath(t, filepath.Join(PowerSupplyPath, "BAT0")) {
		return
	}

	_, err = PowerSupplies("BAT*")
	tErrorIf(t, err)
}
