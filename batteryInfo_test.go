package sstat

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func checkBat(t *testing.T) bool {
	var err error

	_, err = os.Stat(filepath.Join(PowerSupplyPath, "BAT0"))
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}

	if err != nil {
		t.Error(err)
	}

	return true
}

func TestBattery(t *testing.T) {
	var err error

	if !checkBat(t) {
		return
	}

	_, err = Battery("BAT0")
	tErrorIf(t, err)
}

func TestBatteries(t *testing.T) {
	var err error

	if !checkBat(t) {
		return
	}

	_, err = Batteries()
	tErrorIf(t, err)
}
