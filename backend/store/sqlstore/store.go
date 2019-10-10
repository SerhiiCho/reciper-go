package sqlstore

import (
	"database/sql"
	"github.com/SerhiiCho/reciper/backend/store"
	_ "github.com/go-sql-driver/mysql"
)

// Store struct
type Store struct {
	db       *sql.DB
	userRepo *UserRepo
}

// New returns pointer on Store
func New(db *sql.DB) *Store {
	return &Store{db: db}
}

// User returns pointer to user repository
func (store *Store) User() store.UserRepo {
	if store.userRepo != nil {
		return store.userRepo
	}

	store.userRepo = &UserRepo{store: store}

	return store.userRepo
}
