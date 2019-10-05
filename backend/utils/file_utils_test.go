package utils

import "testing"

func TestFileExists(t *testing.T) {
	t.Parallel()

	if !FileExists("file_utils.go") {
		t.Error("FileExists must return true because file_exists_test.go file exists")
	}

	if FileExists("nostradamus") {
		t.Error("FileExists must return false because nostradamus file doesn't exist")
	}

	if FileExists("") {
		t.Error("FileExists must return false because empty string is passed as argument")
	}
}
