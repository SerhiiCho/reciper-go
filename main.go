package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/SerhiiCho/reciper/src/bootstrap"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// HandleError method
func HandleError(text string, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		date := time.Now()

		msgError := "[" + date.Format("02-01-2006 15:04:05") + "] " +
			file + ":" + strconv.Itoa(line) + ": " +
			text + ". " + err.Error() + "\n"

		f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0600)

		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(msgError); err != nil {
			panic(err)
		}
	}
}

func main() {
	// err := errors.New("Some error")
	// HandleError("some error text is here", err)
	bootstrap.IndexPage()
}
