package store

import (
	"github.com/SerhiiCho/reciper/backend/model"
)

// GetRecipes returns all recipes from database
func (store *Store) GetRecipes() []*model.Recipe {
	var recipes []*model.Recipe

	//err := store.Find(&recipes).Error
	//utils.HandleError("Error while getting recipes from database", err, "")

	return recipes
}

// CreateRecipe adds given recipe to a database
func (store *Store) CreateRecipe(recipe *model.Recipe) {
	//err := store.Create(recipe).Error
	//utils.HandleError("Error creating recipe in database", err, "")
}
