package models

import "testing"

func TestCreateRecipe(t *testing.T) {
	t.Parallel()

	recipe := NewRecipe()
	recipe.AddRecipe(Recipe{})

	if len(recipe.Recipes) != 1 {
		t.Error("Recipe item was not created via AddRecipe method")
	}
}

func TestIndexRecipe(t *testing.T) {
	t.Parallel()

	recipe := NewRecipe()
	recipe.AddRecipe(Recipe{})
	result := len(recipe.IndexRecipe())

	if result != 1 {
		t.Errorf("Results must be equal to 1 but got %v", result)
	}
}
