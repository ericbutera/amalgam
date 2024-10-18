package test

import (
	"errors"
	"path"
	"runtime"
)

func GetTestDataPath(datafile string) (string, error) {
	// https://hackandsla.sh/posts/2020-11-23-golang-test-fixtures/
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to discover path")
	}
	return path.Join(path.Dir(filename), "..", "..", "testdata", datafile), nil
}
