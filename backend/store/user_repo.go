package store

import (
	"github.com/SerhiiCho/reciper/backend/model"
)

// UserRepo database repository
type UserRepo struct {
	store *Store
}

// Create creates new user in database
func (repo *UserRepo) Create(user *model.User) (*model.User, error) {
	stmt, stmtErr := repo.store.db.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")

	if stmtErr != nil {
		return nil, stmtErr
	}

	res, execErr := stmt.Exec(user.Email, user.Password)

	if execErr != nil {
		return nil, execErr
	}

	userID, resErr := res.LastInsertId()

	if resErr != nil {
		return nil, resErr
	}

	user.ID = uint(userID)

	return user, nil
}

// FindByEmail returns user from database with given email
func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
