package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "root:111111@tcp(localhost:1001)/reciper_test?charset=utf8"
	}

	os.Exit(m.Run())
}
