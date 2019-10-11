package teststore

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store"
)

// UserRepo struct
type UserRepo struct {
	store *Store
	users map[uint]*model.User
}

// CreateUser creates new user in database
func (repo *UserRepo) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	user.ID = uint(len(repo.users)) + 1
	repo.users[user.ID] = user

	return nil
}

// FindByEmail returns user by given email
func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// FindUser returns user by given id
func (repo *UserRepo) FindUser(id uint) (*model.User, error) {
	user, ok := repo.users[id]

	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return user, nil
}
