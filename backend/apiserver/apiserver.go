package apiserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/SerhiiCho/reciper/backend/store/sqlstore"
	"github.com/gorilla/sessions"
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

	log.Printf("Server started on port %s\n", config.BindAddr)

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
