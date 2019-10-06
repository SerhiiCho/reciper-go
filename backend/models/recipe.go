package models

import (
	"database/sql"
	"github.com/SerhiiCho/reciper/backend/utils"
)

// Recipe model
type Recipe struct {
	User        *User  `json:"user"`
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Excerpt     string `json:"excerpt"`
	Intro       string `json:"intro"`
	Text        string `json:"text"`
	Ingredients string `json:"ingredients"`
	Slug        string `json:"slug"`
	Time        uint64 `json:"time"`
	Image       string `json:"image"`
	Ready       byte   `json:"ready"`
	Approved    byte   `json:"approved"`
	Published   byte   `json:"published"`
	Simple      byte   `json:"simple"`
	CreatedAt   string `json:"created_at"`
}

// RecipeRepo method holds slice of recipes
type RecipeRepo struct {
	DB      *sql.DB
	Recipes []Recipe
}

// NewRecipeRepo method returns pointer to RecipeRepo
func NewRecipeRepo(DB *sql.DB) *RecipeRepo {
	return &RecipeRepo{
		DB:      DB,
		Recipes: []Recipe{},
	}
}

// Length returns the amount of recipe items
func (repo *RecipeRepo) Length() int {
	return len(repo.Recipes)
}

// AddRecipe adds method new recipe
func (repo *RecipeRepo) AddRecipe(recipe Recipe) {
	repo.Recipes = append(repo.Recipes, recipe)
}

// IndexRecipe method returns slice of all recipes
func (repo *RecipeRepo) IndexRecipe() []Recipe {
	if repo.Length() > 0 {
		return repo.Recipes
	}

	rows, err := repo.DB.Query(`
		SELECT id, title_ru, intro_ru, ingredients_ru, text_ru, slug, time, image, ready_ru, approved_ru, published_ru, simple, created_at
		FROM recipes
	`)
	utils.HandleError("Error while getting recipes from database", err)

	for rows.Next() {
		var r Recipe

		scanErr := rows.Scan(&r.ID, &r.Title, &r.Intro, &r.Ingredients, &r.Text, &r.Slug, &r.Time, &r.Image, &r.Ready, &r.Approved, &r.Published, &r.Simple, &r.CreatedAt)
		utils.HandleError("Rows scan error", scanErr)

		repo.AddRecipe(r)
	}

	closeErr := rows.Close()
	utils.HandleError("Closing DB connection error", closeErr)

	return repo.Recipes
}
