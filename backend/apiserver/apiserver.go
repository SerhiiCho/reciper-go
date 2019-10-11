package apiserver

import (
	"database/sql"
	"github.com/SerhiiCho/reciper/backend/store/sqlstore"
	"net/http"
)

// Start
func Start(config *Config) error {
	db, dbErr := newDB(config.DatabaseURL)

	if dbErr != nil {
		return dbErr
	}

	defer db.Close()

	store := sqlstore.New(db)
	serv := newServer(store)

	return http.ListenAndServe(config.BindAddr, serv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, dbErr := sql.Open("mysql", databaseURL)

	if dbErr != nil {
		return nil, dbErr
	}

	if pingErr := db.Ping(); pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}
