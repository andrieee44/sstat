package sstat

import "path/filepath"

// Battery returns battery information in located in [BatteryPath] + basepath.
// basepath can be "BAT0", for example.
func Battery(basepath string) (*PowerSupplyInfo, error) {
	var (
		batteryInfo *PowerSupplyInfo
		err         error
	)

	batteryInfo = new(PowerSupplyInfo)

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

// Batteries returns all system batteries and their information.
func Batteries() ([]*PowerSupplyInfo, error) {
	var (
		batPaths     []string
		batteryInfos []*PowerSupplyInfo
		idx          int
		err          error
	)

	batPaths, err = filepath.Glob(filepath.Join(BatteryPath, "BAT*"))
	if err != nil {
		return nil, err
	}

	batteryInfos = make([]*PowerSupplyInfo, len(batPaths))

	for idx = range batPaths {
		batteryInfos[idx], err = Battery(batPaths[idx])
		if err != nil {
			return nil, err
		}
	}

	return batteryInfos, nil
}
