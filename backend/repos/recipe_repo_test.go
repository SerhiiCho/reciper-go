package repos

import (
	"github.com/SerhiiCho/reciper/backend/models"
	"testing"
)

func TestRecipeRepo_AddRecipe(t *testing.T) {
	t.Parallel()

	recipe := RecipeRepo{}
	recipe.AddRecipe(models.Recipe{})

	if len(recipe.Recipes) != 1 {
		t.Error("Recipe item was not created via AddRecipe method")
	}
}

func TestRecipeRepo_IndexRecipe(t *testing.T) {
	t.Parallel()

	recipe := RecipeRepo{}
	recipe.AddRecipe(models.Recipe{})
	result := len(recipe.IndexRecipe())

	if result != 1 {
		t.Errorf("Results must be equal to 1 but got %v", result)
	}
}

func TestRecipeRepo_Length(t *testing.T) {
	t.Parallel()

	recipe := RecipeRepo{}
	recipe.AddRecipe(models.Recipe{})
	recipe.AddRecipe(models.Recipe{})
	recipe.AddRecipe(models.Recipe{})

	result := recipe.Length()

	if result != 3 {
		t.Errorf("The result must be 3 but got %d", result)
	}
}
