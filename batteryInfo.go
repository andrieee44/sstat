package sstat

// BatteryInfo reports battery information. Documentation for the
// object methods are taken from [sysfs-class-power]. Missing
// documentation or methods means that the author's system
// doesn't have the correct uevent keys or the file
// [sysfs-class-power] itself is missing documentation.
// For missing uevent keys use the method [PowerSupplyInfo.UeventKey].
//
// [sysfs-class-power]: https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux.git/tree/Documentation/ABI/testing/sysfs-class-power?h=v6.14.3
type BatteryInfo struct {
	PowerSupplyInfo
}

// Status reports the charging status of the battery.
// Valid values are "Unknown", "Charging", "Discharging"
// "Not charging" and "Full".
func (info *BatteryInfo) Status() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_STATUS")
}

// Present reports whether a battery is present or not in the system. If the
// property does not exist, the battery is considered to be present.
// Valid values are 0 for absent, 1 for present.
func (info *BatteryInfo) Present() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_PRESENT")
}

// Technology reports the battery technology supported by the supply.
// Valid values are "Unknown", "NiMH", "Li-ion", "Li-poly", "LiFe",
// "NiCd" and "LiMn".
func (info *BatteryInfo) Technology() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_TECHNOLOGY")
}

// CycleCount reports the number of full charge + discharge cycles the
// battery has undergone.
// Valid values are Integer > 0 if representing full cycles or
// Integer = 0 if cycle_count info is not available.
func (info *BatteryInfo) CycleCount() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_CYCLE_COUNT")
}

func (info *BatteryInfo) VoltageMinDesign() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_VOLTAGE_MIN_DESIGN")
}

// VoltageNow Reports an instant, single VBAT voltage reading for the
// battery. This value is not averaged/smoothed.
// Valid values are represented in microvolts.
func (info *BatteryInfo) VoltageNow() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_VOLTAGE_NOW")
}

func (info *BatteryInfo) PowerNow() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_POWER_NOW")
}

func (info *BatteryInfo) EnergyFullDesign() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_ENERGY_FULL_DESIGN")
}

func (info *BatteryInfo) EnergyFull() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_ENERGY_FULL")
}

func (info *BatteryInfo) EnergyNow() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_ENERGY_NOW")
}

// Capacity reports fine grain representation of battery capacity.
// Valid values are 0 - 100 (percent).
func (info *BatteryInfo) Capacity() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_CAPACITY")
}

// CapacityLevel reports coarse representation of battery capacity.
// Valid values are "Unknown", "Critical", "Low", "Normal",
// "High" and "Full".
func (info *BatteryInfo) CapacityLevel() (string, bool) {
	return info.UeventKey("POWER_SUPPLY_CAPACITY_LEVEL")
}

// Battery returns battery information in located in [PowerSupplyPath] + basepath.
func Battery(basepath string) (*BatteryInfo, error) {
	var (
		powerSupplyInfo *PowerSupplyInfo
		err             error
	)

	powerSupplyInfo, err = PowerSupply(basepath)
	if err != nil {
		return nil, err
	}

	return &BatteryInfo{PowerSupplyInfo: *powerSupplyInfo}, nil
}

// Batteries returns all system batteries and their information.
func Batteries() ([]*BatteryInfo, error) {
	var (
		powerSupplyInfos []*PowerSupplyInfo
		batteryInfos     []*BatteryInfo
		idx              int
		err              error
	)

	powerSupplyInfos, err = PowerSupplies("BAT*")
	if err != nil {
		return nil, err
	}

	batteryInfos = make([]*BatteryInfo, len(powerSupplyInfos))

	for idx = range powerSupplyInfos {
		batteryInfos[idx] = &BatteryInfo{PowerSupplyInfo: *powerSupplyInfos[idx]}
	}

	return batteryInfos, nil
}
