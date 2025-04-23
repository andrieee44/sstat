package sstat

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// MemInfo reports memory usage information from /proc/meminfo.
// Documentation for methods are taken from
// [proc_meminfo(5)].
//
// [proc_meminfo(5)]: https://man.archlinux.org/man/proc_meminfo.5.en
type MemInfo struct {
	info map[string]int
}

// Key reports the value of the specified /proc/meminfo parameter
// and whether if the key is valid or not.
func (info *MemInfo) Key(key string) (value int, ok bool) {
	value, ok = info.info[key]

	return value, ok
}

// MemTotal reports total usable RAM (i.e., physical RAM
// minus a few reserved bits and the kernel binary code).
func (info *MemInfo) MemTotal() (value int, ok bool) {
	return info.Key("MemTotal")
}

// MemFree reports the sum of [LowFree]+[HighFree].
func (info *MemInfo) MemFree() (value int, ok bool) {
	return info.Key("MemFree")
}

// MemAvailable (since Linux 3.14) reports the estimate of how much
// memory is available for starting new applications,
// without swapping.
func (info *MemInfo) MemAvailable() (value int, ok bool) {
	return info.Key("MemAvailable")
}

// NewMemInfo returns memory usage information from /proc/meminfo.
func NewMemInfo() (*MemInfo, error) {
	var (
		memInfo *MemInfo
		err     error
	)

	memInfo = &MemInfo{
		info: make(map[string]int),
	}

	err = ScanFile("/proc/meminfo", bufio.ScanLines, func(text string) (bool, error) {
		var (
			fields []string
			value  int
			err    error
		)

		fields = strings.Fields(text)
		if len(fields) != 3 {
			return false, fmt.Errorf("/proc/meminfo: invalid meminfo format")
		}

		value, err = strconv.Atoi(fields[1])
		if err != nil {
			return false, err
		}

		memInfo.info[fields[0][:len(fields[0])-1]] = value

		return true, nil
	})

	return memInfo, err
}
