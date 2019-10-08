package repos

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/storage"
	"testing"
)

func TestRecipeRepo_Add(t *testing.T) {
	t.Parallel()

	recipe := NewRecipeRepo(storage.GetDB())
	recipe.Add(model.Recipe{})

	if len(recipe.Recipes) != 1 {
		t.Error("Recipe item was not created via Add method")
	}
}

func TestRecipeRepo_All(t *testing.T) {
	t.Parallel()
	storage.TestSetup()

	recipe := NewRecipeRepo(storage.GetDB())
	recipe.Add(model.Recipe{})
	result := len(recipe.All())

	if result != 1 {
		t.Errorf("Results must be equal to 1 but got %v", result)
	}
}

func TestRecipeRepo_Length(t *testing.T) {
	t.Parallel()

	recipe := RecipeRepo{}
	recipe.Add(model.Recipe{})
	recipe.Add(model.Recipe{})
	recipe.Add(model.Recipe{})

	result := recipe.Length()

	if result != 3 {
		t.Errorf("The result must be 3 but got %d", result)
	}
}
