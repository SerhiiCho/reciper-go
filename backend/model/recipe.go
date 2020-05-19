package model

import valid "github.com/go-ozzo/ozzo-validation"

// Recipe model
type Recipe struct {
	Model
	User        User   `json:"-"`
	UserID      uint   `json:"-"`
	Title       string `json:"title"`
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
}

// Validate validates recipe data
func (recipe *Recipe) Validate() error {
	return valid.ValidateStruct(
		recipe,
		valid.Field(&recipe.Title, valid.Required),
	)
}
