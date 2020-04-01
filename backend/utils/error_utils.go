package utils

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

// HandleError saves error message to a log file if it is not nil
// returns true if there is an error
func HandleError(text string, err error, filePath string) bool {
	if filePath == "" {
		filePath = "../logs.log"
	}

	if err == nil {
		return false
	}

	_, fileName, lineNum, _ := runtime.Caller(1)
	date := time.Now()

	dateInfo := "[" + date.Format("02-01-2006 15:04:05") + "] "
	msgError := fileName + ":" + strconv.Itoa(lineNum) + ": " + text + ". " + err.Error() + "\n"

	if !FileExists(filePath) {
		FilePutContent(filePath, dateInfo+msgError)
		return true
	}

	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	printIfExist(err)

	log.Println(msgError)

	_, writeErr := logFile.WriteString(dateInfo + msgError)
	printIfExist(writeErr)
	printIfExist(logFile.Close())

	return true
}

// FatalIfError logs the error end exists the program
// if error is not nil
func FatalIfError(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func printIfExist(err error) {
	if err != nil {
		log.Println(err)
	}
}
