package file

import (
	"testing"
)

func TestReadFile1(t *testing.T) {
	v := ReadFile("file_test.txt")
	if v != "file_test" {
		t.Errorf("ReadFile(\"file_test.txt\") failed. Got %v, expected \"file_test\".", v)
	}
}
