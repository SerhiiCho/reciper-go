package sqlstore

import (
	"github.com/SerhiiCho/reciper/backend/model"
)

// GetRecipes returns all recipes from database
func (store *Store) GetRecipes() []*model.Recipe {
	var recipes []*model.Recipe

	return recipes
}

// CreateRecipe adds given recipe to a database
func (store *Store) CreateRecipe(recipe *model.Recipe) {
	//
}
