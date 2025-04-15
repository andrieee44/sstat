package sstat

import (
	"bufio"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// PathReadStr reads the file in located in path. The file is assumed
// to have only one line delimited by a newline. Useful for reading
// battery information from /sys/class/power_supply/BAT0/status for
// example.
func PathReadStr(path string) (string, error) {
	var (
		buf []byte
		err error
	)

	buf, err = os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(buf[:len(buf)-1]), nil
}

// PathReadInt reads the file in located in path and converts the contents of
// the file to an integer. The file is assumed to have only one line
// delimited by a newline. Useful for reading battery information from
// /sys/class/power_supply/BAT0/capacity for example.
func PathReadInt(path string) (int, error) {
	var (
		str string
		num int
		err error
	)

	str, err = PathReadLine(path)
	if err != nil {
		return 0, err
	}

	num, err = strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return num, nil
}

// Meminfo returns a map containing the values in /proc/meminfo.
func Meminfo() (map[string]int, error) {
	var (
		meminfo *os.File
		scanner *bufio.Scanner
		fields  []string
		key     string
		val     int
		err     error
	)

	keyVal = make(map[string]int)

	meminfo, err = os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	scanner = bufio.NewScanner(meminfo)

	for scanner.Scan() {
		fields = strings.Fields(scanner.Text())
		key = fields[0][:len(fields[0])-1]

		keyVal[key], err = strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
	}

	if scanner.Err() != nil {
		return nil, err
	}

	err = meminfo.Close()
	if err != nil {
		return nil, err
	}

	return keyVal, nil
}
