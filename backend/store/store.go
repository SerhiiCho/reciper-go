package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Store struct
type Store struct {
	config   *Config
	db       *sql.DB
	userRepo *UserRepo
}

// NewStore returns pointer on Store
func NewStore(conf *Config) *Store {
	return &Store{config: conf}
}

// Open configures database and returns nil
// If sql error returns error
func (store *Store) Open() error {
	db, err := sql.Open("mysql", store.config.DatabaseURL)

	if err != nil {
		return err
	}

	store.db = db

	return nil
}

// Close closes database connection and returns nil
// If failed closing database returns error
func (store *Store) Close() error {
	if err := store.db.Close(); err != nil {
		return err
	}

	return nil
}

// User returns pointer to user repository
func (store *Store) User() *UserRepo {
	if store.userRepo != nil {
		return store.userRepo
	}

	store.userRepo = &UserRepo{store: store}

	return store.userRepo
}
