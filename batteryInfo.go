package sstat

//
type BatteryInfo struct {
	PowerSupplyInfo
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
