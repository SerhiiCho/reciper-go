package apiserver

import (
	"database/sql"
	"github.com/SerhiiCho/reciper/backend/store/sqlstore"
	"github.com/gorilla/sessions"
	"net/http"
)

// Start starts the api server
func Start(config *Config) error {
	db, dbErr := newDB(config.DatabaseURL)

	if dbErr != nil {
		return dbErr
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	server := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, server)
}

// newDB configures the database connection
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
