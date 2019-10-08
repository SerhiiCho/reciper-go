package db

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/utils"
)

func (db *Database) GetRecipes() []*model.Recipe {
	var recipes []*model.Recipe

	err := db.Find(&recipes).Error
	utils.HandleError("Error while getting recipes from database", err, "")

	return recipes
}
