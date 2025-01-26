package test

import (
	"errors"
	"io"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

var ErrInvalidPath = errors.New("invalid path")

func GetTestDataPath(datafile string) (string, error) {
	// https://hackandsla.sh/posts/2020-11-23-golang-test-fixtures/
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", ErrInvalidPath
	}
	return path.Join(path.Dir(filename), "..", "..", "testdata", datafile), nil
}

func FileToReadCloser(file string) (io.ReadCloser, error) {
	path, err := GetTestDataPath(file)
	if err != nil {
		return nil, err
	}
	return os.Open(path)
}

// diff compares two structs and prints the differences
func Diff(t *testing.T, expected any, actual any, ignoredFields ...string) {
	// https://github.com/google/go-cmp/blob/391980c4b2e1cc2c30d2bfae6039815350490495/cmp/cmpopts/example_test.go#L32-L34
	t.Helper()
	ignored := cmpopts.IgnoreFields(expected, ignoredFields...)
	if d := cmp.Diff(expected, actual, ignored); d != "" {
		t.Errorf("mismatch (-want +got):\n%s", d)
	}
}

func DiffProto(t *testing.T, expected proto.Message, actual proto.Message) {
	t.Helper()
	diff := cmp.Diff(expected, actual, protocmp.Transform())
	if diff != "" {
		t.Errorf("unexpected diff: %s", diff)
	}
}
