package recipe

import "github.com/SerhiiCho/reciper/backend/models/user"

// Recipe model
type Model struct {
	User        user.Model `json:"user"`
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Excerpt     string     `json:"excerpt"`
	Intro       string     `json:"intro"`
	Text        string     `json:"text"`
	Ingredients string     `json:"ingredients"`
	Slug        string     `json:"slug"`
	Time        uint64     `json:"time"`
	Image       string     `json:"image"`
	Ready       bool       `json:"ready"`
	Approved    bool       `json:"approved"`
	Published   bool       `json:"published"`
	Simple      bool       `json:"simple"`
	CreatedAt   string     `json:"created_at"`
}

// Repo method holds slice of recipes
type Repo struct {
	Recipes []Model
}

// New method returns pointer to Repo
func New() *Repo {
	return &Repo{
		Recipes: []Model{},
	}
}

// Adds method new recipe
func (r *Repo) Add(recipe Model) {
	r.Recipes = append(r.Recipes, recipe)
}

// GetAll method returns slice of all recipes
func (r *Repo) GetAll() []Model {
	return r.Recipes
}
