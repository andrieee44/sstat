package sstat

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

func checkPath(t *testing.T, path string) bool {
	var err error

	_, err = os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}

	if err != nil {
		t.Error(err)
	}

	return true
}

func tErrorIf(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func tmpFile(content string) (string, error) {
	var (
		file *os.File
		err  error
	)

	file, err = os.CreateTemp(os.TempDir(), "sstat-test-")
	if err != nil {
		return "", err
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func TestPathReadStr(t *testing.T) {
	var (
		path, value string
		err         error
	)

	path, err = tmpFile("hello\n")
	tErrorIf(t, err)

	value, err = PathReadStr(path)
	tErrorIf(t, err)
	tErrorIf(t, os.Remove(path))

	if value != "hello" {
		t.Errorf("expected %q, got %q", "hello", value)
	}
}

func TestPathReadInt(t *testing.T) {
	var (
		path  string
		value int
		err   error
	)

	path, err = tmpFile("123\n")
	tErrorIf(t, err)

	value, err = PathReadInt(path)
	tErrorIf(t, err)
	tErrorIf(t, os.Remove(path))

	if value != 123 {
		t.Errorf("expected %d, got %d", 123, value)
	}
}
