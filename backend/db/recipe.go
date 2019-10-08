package db

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/utils"
)

// GetRecipes returns all recipes from database
func (db *Database) GetRecipes() []*model.Recipe {
	var recipes []*model.Recipe

	err := db.Find(&recipes).Error
	utils.HandleError("Error while getting recipes from database", err, "")

	return recipes
}

// CreateRecipe adds given recipe to a database
func (db *Database) CreateRecipe(recipe *model.Recipe) {
	err := db.Create(recipe).Error
	utils.HandleError("Error creating recipe in database", err, "")
}
