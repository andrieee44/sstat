package sstat

import (
	"path/filepath"
	"testing"
)

func TestBattery(t *testing.T) {
	var err error

	if !checkPath(t, filepath.Join(PowerSupplyPath, "BAT0")) {
		return
	}

	_, err = Battery("BAT0")
	tErrorIf(t, err)
}

func TestBatteries(t *testing.T) {
	var err error

	if !checkPath(t, filepath.Join(PowerSupplyPath, "BAT0")) {
		return
	}

	_, err = Batteries()
	tErrorIf(t, err)
}
