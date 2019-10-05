package models

// Recipe model
type Recipe struct {
	User        User   `json:"user"`
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Excerpt     string `json:"excerpt"`
	Intro       string `json:"intro"`
	Text        string `json:"text"`
	Ingredients string `json:"ingredients"`
	Slug        string `json:"slug"`
	Time        uint   `json:"time"`
	Image       string `json:"image"`
	Ready       bool   `json:"ready"`
	Approved    bool   `json:"approved"`
	Published   bool   `json:"published"`
	Simple      bool   `json:"simple"`
	CreatedAt   string `json:"created_at"`
}
