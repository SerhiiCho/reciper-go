package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Store struct
type Store struct {
	config *Config
	db     *sql.DB
}

// NewStore opens database
func NewStore(conf *Config) *Store {
	return &Store{
		config: conf,
	}
}

func (store *Store) Open() error {
	db, err := sql.Open("mysql", store.config.DatabaseURL)

	if err != nil {
		return err
	}

	store.db = db

	return nil
}

func (store *Store) Close() error {
	if err := store.db.Close(); err != nil {
		return err
	}

	return nil
}
