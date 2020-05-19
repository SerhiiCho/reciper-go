package model

import valid "github.com/go-ozzo/ozzo-validation"

// Recipe model
type Recipe struct {
	Model
	User          User   `json:"-"`
	UserID        uint   `json:"-"`
	TitleRu       string `json:"title_ru"`
	TitleEn       string `json:"title_en"`
	IntroRu       string `json:"intro_ru"`
	IntroEn       string `json:"intro_en"`
	TextRu        string `json:"text_ru"`
	TextEn        string `json:"text_en"`
	IngredientsRu string `json:"ingredients_ru"`
	IngredientsEn string `json:"ingredients_en"`
	Slug          string `json:"slug"`
	Time          uint64 `json:"time"`
	Image         string `json:"image"`
	ReadyRu       byte   `json:"ready_ru"`
	ReadyEn       byte   `json:"ready_en"`
	ApprovedRu    byte   `json:"approved_ru"`
	ApprovedEn    byte   `json:"approved_en"`
	PublishedRu   byte   `json:"published_ru"`
	PublishedEn   byte   `json:"published_en"`
	Simple        byte   `json:"simple"`
}

// Validate validates recipe data
func (recipe *Recipe) Validate() error {
	return valid.ValidateStruct(
		recipe,
		valid.Field(&recipe.TitleRu, valid.Required),
	)
}
