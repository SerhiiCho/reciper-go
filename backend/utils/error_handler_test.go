package utils

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"
	"time"
)

func TestHandleError(t *testing.T) {
	t.Parallel()
}

func TestPrintIfExist(t *testing.T) {
	t.Run("if argument is nil", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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
