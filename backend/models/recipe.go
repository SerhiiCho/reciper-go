package models

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
	Ready       bool   `json:"ready"`
	Approved    bool   `json:"approved"`
	Published   bool   `json:"published"`
	Simple      bool   `json:"simple"`
	CreatedAt   string `json:"created_at"`
}

// RecipeRepo method holds slice of recipes
type RecipeRepo struct {
	Recipes []Recipe
}

// NewRecipe method returns pointer to RecipeRepo
func NewRecipe() *RecipeRepo {
	return &RecipeRepo{
		Recipes: []Recipe{},
	}
}

// CreateRecipe adds method new recipe
func (r *RecipeRepo) CreateRecipe(recipe Recipe) {
	r.Recipes = append(r.Recipes, recipe)
}

// IndexRecipe method returns slice of all recipes
func (r *RecipeRepo) IndexRecipe() []Recipe {
	return r.Recipes
}
