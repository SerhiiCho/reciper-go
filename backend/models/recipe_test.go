package models

import "testing"

func TestCreateRecipe(t *testing.T) {
	t.Parallel()
	recipe := NewRecipe()
	recipe.CreateRecipe(Recipe{})

	if len(recipe.Recipes) != 1 {
		t.Error("Item was not added")
	}
}

func TestIndexRecipe(t *testing.T) {
	t.Parallel()

	recipe := NewRecipe()
	recipe.CreateRecipe(Recipe{})
	result := len(recipe.IndexRecipe())

	if result != 1 {
		t.Errorf("Results must be equal to 1 but got %v", result)
	}
}
