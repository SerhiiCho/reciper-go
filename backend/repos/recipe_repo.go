package repos

import (
	"database/sql"
	"github.com/SerhiiCho/reciper/backend/models"
	"github.com/SerhiiCho/reciper/backend/utils"
)

// RecipeRepo method holds slice of recipes
type RecipeRepo struct {
	DB      *sql.DB
	Recipes []models.Recipe
}

// NewRecipeRepo method returns pointer to RecipeRepo
func NewRecipeRepo(DB *sql.DB) *RecipeRepo {
	return &RecipeRepo{
		DB:      DB,
		Recipes: []models.Recipe{},
	}
}

// Length returns the amount of recipe items
func (repo *RecipeRepo) Length() int {
	return len(repo.Recipes)
}

// AddRecipe adds method new recipe
func (repo *RecipeRepo) AddRecipe(recipe models.Recipe) {
	repo.Recipes = append(repo.Recipes, recipe)
}

// IndexRecipe method returns slice of all recipes
func (repo *RecipeRepo) IndexRecipe() []models.Recipe {
	if repo.Length() > 0 {
		return repo.Recipes
	}

	rows, err := repo.DB.Query(`
		SELECT id, title_ru, intro_ru, ingredients_ru, text_ru, slug, time, image, ready_ru, approved_ru, published_ru, simple, created_at
		FROM recipes
	`)
	utils.HandleError("Error while getting recipes from database", err)

	for rows.Next() {
		var r models.Recipe

		scanErr := rows.Scan(&r.ID, &r.Title, &r.Intro, &r.Ingredients, &r.Text, &r.Slug, &r.Time, &r.Image, &r.Ready, &r.Approved, &r.Published, &r.Simple, &r.CreatedAt)
		utils.HandleError("Rows scan error", scanErr)

		repo.AddRecipe(r)
	}

	closeErr := rows.Close()
	utils.HandleError("Closing DB connection error", closeErr)

	return repo.Recipes
}
