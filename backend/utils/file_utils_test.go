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

func TestFileGetContents(t *testing.T) {
	t.Parallel()

	t.Run("returns file content", func(t *testing.T) {
		FilePutContent("test-file", "Hello world")
		result := FileGetContent("test-file")

		if result != "Hello world" {
			t.Errorf("The result must be `Hello world`%v", result)
		}
		RemoveFileIfExist("test-file")
	})

	t.Run("returns empty string if file is empty", func(t *testing.T) {
		FilePutContent("test-file", "")
		result := FileGetContent("")

		if result != "" {
			t.Error("The value of result variable must be empty string")
		}
		RemoveFileIfExist("test-file")
	})

	t.Run("returns empty string if does not exist", func(t *testing.T) {
		result := FileGetContent("nostradamus")

		if result != "" {
			t.Error("FileGetContents must return empty string because file doesn't exist")
		}
	})
}
