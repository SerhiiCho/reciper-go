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
		FilePutContent("nice-file", "Hello world")
		result := FileGetContent("nice-file")

		if result != "Hello world" {
			t.Errorf("The result must be `Hello world`%v", result)
		}
		RemoveFileIfExist("nice-file")
	})

	t.Run("returns empty string if file is empty", func(t *testing.T) {
		FilePutContent("testing-file", "")

		if FileGetContent("") != "" {
			t.Error("The value of result variable must be empty string")
		}

		RemoveFileIfExist("testing-file")
	})

	t.Run("returns empty string if does not exist", func(t *testing.T) {
		result := FileGetContent("nostradamus")

		if result != "" {
			t.Error("FileGetContents must return empty string because file doesn't exist")
		}
	})
}

func TestFilePutContent(t *testing.T) {
	t.Parallel()

	t.Run("creates file if it doesn't exist", func(t *testing.T) {
		RemoveFileIfExist("test-file")

		FilePutContent("test-file", "")

		if !FileExists("test-file") {
			t.Error("test-file doesn't exist after creation")
		}
	})

	t.Run("inserts text in a file", func(t *testing.T) {
		FilePutContent("test-file", "Hello planet!")

		if FileGetContent("test-file") != "Hello planet!" {
			t.Error("test-file doesn't contain text `Hello planet!` after inserting it")
		}

		RemoveFileIfExist("test-file")
	})
}
