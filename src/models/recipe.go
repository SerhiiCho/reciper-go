package models

// Recipes model
type Recipe struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Intro string `json:"intro,omitempty"`
	Text  string `json:"text,omitempty"`
}
