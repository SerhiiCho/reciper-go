package models

import "testing"

func TestCreateRecipe(t *testing.T) {
	recipe := NewRecipe()
	recipe.CreateRecipe(Recipe{})

	if len(recipe.Recipes) != 1 {
		t.Error("Item was not added")
	}
}

func TestIndexRecipe(t *testing.T) {
	recipe := NewRecipe()
	recipe.CreateRecipe(Recipe{})

	if len(recipe.IndexRecipe()) != 1 {
		t.Error("Results must be equal to 1")
	}
}
