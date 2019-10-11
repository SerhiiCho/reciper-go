package teststore

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store"
)

// UserRepo struct
type UserRepo struct {
	store *Store
	users map[string]*model.User
}

// CreateUser creates new user in database
func (repo *UserRepo) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	repo.users[user.Email] = user
	user.ID = uint(len(repo.users))

	return nil
}

// FindByEmail returns user given email
func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	user, ok := repo.users[email]

	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return user, nil
}
