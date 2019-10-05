package utils

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestHandleError(t *testing.T) {
	t.Parallel()

	t.Run("creates file if it doesnt exist", func(t *testing.T) {
		RemoveFileIfExist("../logs.log")

		customError := errors.New("my error message")
		HandleError("Some message", customError)

		logContent := FileGetContent("../logs.log")

		if !strings.Contains(logContent, "my error message") {
			t.Error("Log file myst contain message: `my error message` but it doesn't")
		}

		RemoveFileIfExist("../logs.log")
	})
}

func TestPrintIfExist(t *testing.T) {
	t.Parallel()

	t.Run("if argument is nil", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)

		defer func() {
			log.SetOutput(os.Stderr)
		}()

		printIfExist(nil)

		result := buf.String()

		if result != "" {
			t.Errorf("Result must be empty string but got %v", result)
		}
	})

	t.Run("if argument is error", func(t *testing.T) {
		var buf bytes.Buffer

		log.SetOutput(&buf)

		defer func() {
			log.SetOutput(os.Stderr)
		}()

		printIfExist(errors.New("some error message"))

		result := buf.String()
		dateTime := time.Now().Format("2006/01/02 15:04:05")
		expected := dateTime + " some error message\n"

		if result != expected {
			t.Errorf("Result must be `%s` but got `%s`", expected, result)
		}
	})
}
