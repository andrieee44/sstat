package sstat

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// BatteryPath is the path where the system battery information is located.
const BatteryPath string = "/sys/class/power_supply"

// BatteryInfo contains battery information.
type BatteryInfo struct {
	// Status is the status of the battery.
	// Possible values are "Charging", "Discharging", "Full", "Not Charging"
	// and "Unknown".
	Status   string
	
	// Capacity is the percentage of the battery.
	Capacity int
}

// Battery returns battery information in located in [BatteryPath] + basepath.
// basepath can be "BAT0", for example.
func Battery(basepath string) (*BatteryInfo, error) {
	var (
		status   []byte
		capacity int
		err      error
	)

	status, err = PathReadStr(filepath.Join(BatteryPath, basepath, "status"))
	if err != nil {
		return nil, err
	}

	capacity, err = PathReadInt(filepath.Join(BatteryPath, basepath, "capacity"))
	if err != nil {
		return nil, err
	}

	return &BatteryInfo{
		Status:   status,
		Capacity: capacity,
	}, nil
}

// Batteries returns all system batteries and their information.
func Batteries() ([]*BatteryInfo, error) {
	var (
		batPaths   []string
		batInfoMap map[string]*BatteryInfo
		path       string
		err        error
	)

	batPaths, err = filepath.Glob(filepath.Join(BatteryPath, "BAT*"))
	if err != nil {
		return nil, err
	}

	batInfoMap = make(map[string]*BatteryInfo)

	for _, path = range batPaths {
		batInfoMap[filepath.Base(path)], err = Battery(path)
		if err != nil {
			return nil, err
		}
	}

	return batInfoMap
}
