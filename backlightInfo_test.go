package sstat

import (
	"path/filepath"
	"testing"
)

func TestBacklight(t *testing.T) {
	var err error

	if !checkPath(t, filepath.Join(BacklightPath, "intel_backlight")) {
		return
	}

	_, err = Backlight("intel_backlight")
	tErrorIf(t, err)
}

func TestBacklights(t *testing.T) {
	var err error

	if !checkPath(t, filepath.Join(BacklightPath, "intel_backlight")) {
		return
	}

	_, err = Backlights("*")
	tErrorIf(t, err)
}
