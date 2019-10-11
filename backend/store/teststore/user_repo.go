package teststore

import (
	"errors"
	"github.com/SerhiiCho/reciper/backend/model"
)

// UserRepo struct
type UserRepo struct {
	store *Store
	users map[string]*model.User
}

// Create creates new user in database
func (repo *UserRepo) Create(user *model.User) error {
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
		return nil, errors.New("user not found")
	}

	return user, nil
}
