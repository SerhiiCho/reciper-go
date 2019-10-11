package teststore

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store"
)

// Store struct
type Store struct {
	userRepo *UserRepo
}

// New returns pointer on Store
func New() *Store {
	return &Store{}
}

// User returns pointer to user repository
func (store *Store) User() store.UserRepo {
	if store.userRepo != nil {
		return store.userRepo
	}

	store.userRepo = &UserRepo{
		store: store,
		users: make(map[string]*model.User),
	}

	return store.userRepo
}
