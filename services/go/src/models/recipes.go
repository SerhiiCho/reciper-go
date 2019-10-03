package models

// Recipes model
type Recipes struct {
	ID      uint   `json:"id,omitempty"`
	TitleRu string `json:"title_ru,omitempty"`
	TitleEn string `json:"title_en,omitempty"`
	IntroRu string `json:"intro_ru,omitempty"`
	IntroEn string `json:"intro_en,omitempty"`
	TextRu  string `json:"text_ru,omitempty"`
	TextEn  string `json:"text_en,omitempty"`
}
