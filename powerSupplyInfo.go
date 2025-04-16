package sstat

// PowerSupplyPath is the directory where the information for
// system power supplies are located.
const PowerSupplyPath string = "/sys/class/power_supply"

// PowerSupplyInfo reports battery information.
type PowerSupplyInfo struct {
	validKeys    map[string]bool
	manufacturer string
	modelName    string
	name         string
	psType       string
	serialNumber string
}

func (info *PowerSupplyInfo) validKey(key string) bool {
	var ok bool

	_, ok = info.validKeys[key]

	return ok
}

// Manufacturer reports the name of the device manufacturer.
func (info *PowerSupplyInfo) Manufacturer() (string, bool) {
	return info.manufacturer, info.validKey("POWER_SUPPLY_MANUFACTURER")
}

// ModelName reports the name of the device model.
func (info *PowerSupplyInfo) ModelName() (string, bool) {
	return info.modelName, info.validKey("POWER_SUPPLY_MODEL_NAME")
}

// SerialNumber reports the serial number of the device.
func (info *PowerSupplyInfo) SerialNumber() (string, bool) {
	return info.serialNumber, info.validKey("POWER_SUPPLY_SERIAL_NUMBER")
}

// Type reports the main type of the supply.
// Valid values are "Battery", "UPS", "Mains", "USB" and "Wireless".
func (info *PowerSupplyInfo) Type() (string, bool) {
	return info.psType, info.validKey("POWER_SUPPLY_TYPE")
}

// Name reports the name of the device.
func (info *PowerSupplyInfo) Name() (string, bool) {
	return info.name, info.validKey("POWER_SUPPLY_NAME")
}

// PowerSupply returns battery information in located in
// [BatteryPath] + basepath.
// basepath can be "BAT0", for example.
func PowerSupply(basepath string) (*PowerSupplyInfo, error) {
	var (
		powerSupplyInfo *PowerSupplyInfo
		err             error
	)

	batteryInfo = new(BatteryInfo)

	batteryInfo.Status, err = PathReadStr(filepath.Join(BatteryPath, basepath, "status"))
	if err != nil {
		return nil, err
	}

	batteryInfo.Capacity, err = PathReadInt(filepath.Join(BatteryPath, basepath, "capacity"))
	if err != nil {
		return nil, err
	}

	return batteryInfo, nil
}

// PowerSupplies returns all system batteries and their information.
func PowerSupplies() ([]*BatteryInfo, error) {
	var (
		batPaths     []string
		batteryInfos []*BatteryInfo
		idx          int
		err          error
	)

	batPaths, err = filepath.Glob(filepath.Join(BatteryPath, "BAT*"))
	if err != nil {
		return nil, err
	}

	batteryInfos = make([]*BatteryInfo, len(batPaths))

	for idx = range batPaths {
		batteryInfos[idx], err = Battery(batPaths[idx])
		if err != nil {
			return nil, err
		}
	}

	return batteryInfos, nil
}
