package utils

import "testing"

func TestFileExists(t *testing.T) {
	t.Parallel()

	if !FileExists("file_exists_test.go") {
		t.Error("FileExists must return true because file_exists_test.go file exists")
	}

	if FileExists("sadlfsdl") {
		t.Error("FileExists must return false because sadlfsdl file doen't exist")
	}

	if FileExists("") {
		t.Error("FileExists must return false because empty string is passed as argument")
	}
}
