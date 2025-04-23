package sstat

import (
	"bufio"
	"os"
	"strconv"
)

// PathReadStr reads the file in located in path. The file is assumed
// to have only one line delimited by a newline.
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
// delimited by a newline.
func PathReadInt(path string) (int, error) {
	var (
		str string
		num int
		err error
	)

	str, err = PathReadStr(path)
	if err != nil {
		return 0, err
	}

	num, err = strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return num, nil
}

// Scan opens the file in path, splitting the contents
// of the file depending on the split function.
// The split text is then sent to the parse function
// with ok == false stopping the further splitting.
//
// See the source code of [NewMemInfo] for an example usage.
func ScanFile(path string, split bufio.SplitFunc, parser func(text string) (ok bool, err error)) error {
	var (
		file    *os.File
		scanner *bufio.Scanner
		ok      bool
		err     error
	)

	file, err = os.Open(path)
	if err != nil {
		return err
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(split)

	for scanner.Scan() {
		ok, err = parser(scanner.Text())
		if err != nil {
			return err
		}

		if !ok {
			break
		}
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return file.Close()
}
