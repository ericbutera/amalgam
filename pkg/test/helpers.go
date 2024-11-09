package test

import (
	"errors"
	"path"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func GetTestDataPath(datafile string) (string, error) {
	// https://hackandsla.sh/posts/2020-11-23-golang-test-fixtures/
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to discover path")
	}
	return path.Join(path.Dir(filename), "..", "..", "testdata", datafile), nil
}

// diff compares two structs and prints the differences
func Diff(t *testing.T, expected any, actual any, ignoredFields ...string) {
	// https://github.com/google/go-cmp/blob/391980c4b2e1cc2c30d2bfae6039815350490495/cmp/cmpopts/example_test.go#L32-L34
	// TODO: research creating a testify assertion
	ignored := cmpopts.IgnoreFields(expected, ignoredFields...)
	if d := cmp.Diff(expected, actual, ignored); d != "" {
		t.Errorf("mismatch (-want +got):\n%s", d)
	}
}
