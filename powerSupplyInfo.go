package sstat

import (
	"bufio"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// PowerSupplyPath is the directory where the information for
// power supply devices are located.
const PowerSupplyPath string = "/sys/class/power_supply"

// PowerSupplyInfo reports power supply device information.
type PowerSupplyInfo struct {
	info map[string]string
}

// Populate sets the values of every integer
// pointer associated with a key.
func (info *PowerSupplyInfo) Populate(vars map[string]*string) error {
	var (
		key     string
		value   *string
		missing []string
		ok      bool
	)

	for key, value = range vars {
		*value, ok = info.Key(key)
		if !ok {
			missing = append(missing, key)
		}
	}

	if len(missing) != 0 {
		return fmt.Errorf("%s: missing PowerSupplyInfo uevent key(s)", strings.Join(missing, ", "))
	}

	return nil
}

// Key reports the value of the specified uevent key
// and whether if the key is valid or not.
func (info *PowerSupplyInfo) Key(key string) (value string, ok bool) {
	value, ok = info.info[key]

	return value, ok
}

// Manufacturer reports the name of the device manufacturer.
func (info *PowerSupplyInfo) Manufacturer() (value string, ok bool) {
	return info.Key("POWER_SUPPLY_MANUFACTURER")
}

// ModelName reports the name of the device model.
func (info *PowerSupplyInfo) ModelName() (value string, ok bool) {
	return info.Key("POWER_SUPPLY_MODEL_NAME")
}

// SerialNumber reports the serial number of the device.
func (info *PowerSupplyInfo) SerialNumber() (value string, ok bool) {
	return info.Key("POWER_SUPPLY_SERIAL_NUMBER")
}

// Type reports the main type of the supply.
//
// Valid values are:
//   - "Battery"
//   - "UPS"
//   - "Mains"
//   - "USB"
//   - "Wireless"
func (info *PowerSupplyInfo) Type() (value string, ok bool) {
	return info.Key("POWER_SUPPLY_TYPE")
}

// Name reports the name of the device.
func (info *PowerSupplyInfo) Name() (value string, ok bool) {
	return info.Key("POWER_SUPPLY_NAME")
}

// PowerSupply returns power supply device information in
// [PowerSupplyPath] + basepath.
func PowerSupply(basepath string) (*PowerSupplyInfo, error) {
	var (
		powerSupplyInfo *PowerSupplyInfo
		err             error
	)

	powerSupplyInfo = &PowerSupplyInfo{
		info: make(map[string]string),
	}

	err = ScanFile(filepath.Join(PowerSupplyPath, basepath, "uevent"), bufio.ScanLines, func(text string) (bool, error) {
		var fields []string

		fields = strings.Split(text, "=")
		if len(fields) != 2 {
			return false, errors.New("invalid uevent format")
		}

		powerSupplyInfo.info[fields[0]] = fields[1]

		return true, nil
	})

	return powerSupplyInfo, err
}

// PowerSupplies returns all power supply device information in
// [PowerSupplyPath] + glob.
func PowerSupplies(glob string) ([]*PowerSupplyInfo, error) {
	var (
		powerSupplyPaths []string
		powerSupplyInfos []*PowerSupplyInfo
		idx              int
		err              error
	)

	powerSupplyPaths, err = filepath.Glob(filepath.Join(PowerSupplyPath, glob))
	if err != nil {
		return nil, err
	}

	powerSupplyInfos = make([]*PowerSupplyInfo, len(powerSupplyPaths))

	for idx = range powerSupplyPaths {
		powerSupplyInfos[idx], err = PowerSupply(filepath.Base(powerSupplyPaths[idx]))
		if err != nil {
			return nil, err
		}
	}

	return powerSupplyInfos, nil
}
