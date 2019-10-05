package utils

import "os"

// RemoveFileIfExist deletes file if it exist
func RemoveFileIfExist(fileName string) {
	if FileExists(fileName) {
		err := os.Remove(fileName)
		HandleError("Error while trying to delete file", err)
	}
}
