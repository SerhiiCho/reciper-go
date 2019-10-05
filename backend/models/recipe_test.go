package models

import "testing"

func TestAdd(t *testing.T) {
	recipe := New()
	recipe.Add(Recipe{})

	if len(recipe.Recipes) != 1 {
		t.Error("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	recipe := New()
	recipe.Add(Recipe{})

	if len(recipe.GetAll()) != 1 {
		t.Error("Results must be equal to 1")
	}
}
