package store

import "github.com/SerhiiCho/reciper/backend/model"

// UserRepo interface
type UserRepo interface {
	CreateUser(*model.User) error
	FindByEmail(string) (*model.User, error)
}

// RecipeRepo interface
type RecipeRepo interface {
	CreateRecipe(*model.Recipe) error
}
