package utils

import (
	"os"
	"runtime"
	"strconv"
	"time"
)

// HandleError saves error message to a log file if it is not nil
// returns true if there is an error
func HandleError(text string, err error) bool {
	if err == nil {
		return false
	}

	_, fileName, lineNum, _ := runtime.Caller(1)
	date := time.Now()

	msgError := "[" + date.Format("02-01-2006 15:04:05") + "] " +
		fileName + ":" + strconv.Itoa(lineNum) + ": " +
		text + ". " + err.Error() + "\n"

	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(msgError); err != nil {
		panic(err)
	}

	closeErr := f.Close()

	if closeErr != nil {
		panic(err)
	}

	return true
}
