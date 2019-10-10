package store

import "github.com/SerhiiCho/reciper/backend/model"

// UserRepo interface
type UserRepo interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
