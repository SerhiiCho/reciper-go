package store

import (
	"testing"
)

// TestStore truncates given tables
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	conf := NewConfig()
	conf.DatabaseURL = databaseURL

	st := NewStore(conf)

	if err := st.Open(); err != nil {
		t.Error(err)
	}

	return st, func(tables ...string) {
		if len(tables) == 0 {
			return
		}

		for _, table := range tables {
			_, execErr := st.db.Exec("TRUNCATE TABLE " + table)

			if execErr != nil {
				t.Fatal(execErr)
			}
		}

		closeErr := st.Close()

		if closeErr != nil {
			t.Fatal(closeErr)
		}
	}
}
