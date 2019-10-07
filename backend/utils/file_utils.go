package utils

import (
	"io/ioutil"
	"os"
)

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// RemoveFileIfExist deletes file if it exist
func RemoveFileIfExist(fileName string) {
	if FileExists(fileName) {
		err := os.Remove(fileName)
		HandleError("Error while trying to delete file", err, "")
	}
}

// FilePutContent creates file and inserts text into it
func FilePutContent(filePath string, text string) {
	err := ioutil.WriteFile(filePath, []byte(text), 0600)
	HandleError("Can't create file in path "+filePath, err, "")
}

// FileGetContent returns the content of the given file
func FileGetContent(filePath string) string {
	if !FileExists(filePath) {
		return ""
	}

	fileText, err := ioutil.ReadFile(filePath)
	HandleError("File reading error", err, "")

	return string(fileText)
}
