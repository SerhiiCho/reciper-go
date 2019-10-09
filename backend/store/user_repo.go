package store

import "github.com/SerhiiCho/reciper/backend/model"

// UserRepo database repository
type UserRepo struct {
	store *Store
}

// Create creates new user in database
func (repo *UserRepo) Create(user *model.User) (*model.User, error) {
	return nil, nil
}

// FindByEmail returns user from database with given email
func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
