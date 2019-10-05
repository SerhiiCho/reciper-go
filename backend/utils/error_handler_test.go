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
	t.Error("empty test")
}

func TestCreateLogFileIfDoesntExist(t *testing.T) {
	t.Error("empty test")
}

func TestPrintIfExist_IfArgumentIsNil(t *testing.T) {
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
}

func TestPrintIfExist_IfArgumentIsError(t *testing.T) {
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
}
