package model

// Recipe model
type Recipe struct {
	Model
	User          User   `json:"-" gorm:"association_foreignkey:UserID"`
	UserID        uint   `json:"-" gorm:"type:int(10);unsigned;not null"`
	TitleRu       string `json:"title_ru" gorm:"type:varchar(100);default:null"`
	TitleEn       string `json:"title_en" gorm:"type:varchar(100);default:null"`
	IntroRu       string `json:"intro_ru" gorm:"type:varchar(210);default:null"`
	IntroEn       string `json:"intro_en" gorm:"type:varchar(210);default:null"`
	TextRu        string `json:"text_ru" gorm:"type:text;default:null"`
	TextEn        string `json:"text_en" gorm:"type:text;default:null"`
	IngredientsRu string `json:"ingredients_ru" gorm:"type:text;default:null"`
	IngredientsEn string `json:"ingredients_en" gorm:"type:text;default:null"`
	Slug          string `json:"slug" gorm:"type:varchar(100);not null"`
	Time          uint64 `json:"time" gorm:"type:smallint(5);unsigned;default:'0'"`
	Image         string `json:"image" gorm:"type:varchar(200);unsigned;default:'default.jpg'"`
	ReadyRu       byte   `json:"ready_ru" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	ReadyEn       byte   `json:"ready_en" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	ApprovedRu    byte   `json:"approved_ru" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	ApprovedEn    byte   `json:"approved_en" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	PublishedRu   byte   `json:"published_ru" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	PublishedEn   byte   `json:"published_en" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	Simple        byte   `json:"simple" form:"type:tinyint(1);unsigned;default:'0';not null"`
}
