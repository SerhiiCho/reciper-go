package model

import "github.com/SerhiiCho/reciper/backend/utils"

// Recipe model
type Recipe struct {
	Model
	User        User   `json:"-"`
	UserID      uint   `json:"-"`
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
}

// GetExcerpt returns the short version of the title
func (recipe *Recipe) GetExcerpt() string {
	return utils.StrLimit(recipe.Title, 40)
}
