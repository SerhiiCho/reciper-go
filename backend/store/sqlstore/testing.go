package sqlstore

import (
	"database/sql"
	"github.com/SerhiiCho/reciper/backend/utils"
	"testing"
)

// TestDB truncates given tables
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, openErr := sql.Open("mysql", databaseURL)
	utils.FatalIfError("Open database error in store/testing@TestDB", openErr)

	return db, func(tables ...string) {
		if len(tables) == 0 {
			return
		}

		for _, table := range tables {
			_, execErr := db.Exec("TRUNCATE TABLE " + table)
			utils.FatalIfError("Exec db error in store/testing@TestDB", execErr)
		}

		utils.FatalIfError("Close db error in store/testing@TestDB", db.Close())
	}
}
