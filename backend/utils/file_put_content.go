package utils

import "io/ioutil"

// FilePutContent creates file and inserts text into it
func FilePutContent(filePath string, text string) {
	err := ioutil.WriteFile(filePath, []byte(text), 0600)
	HandleError("Can't create file in path "+filePath, err)
}
