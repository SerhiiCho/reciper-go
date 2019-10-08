package model

// Recipe model
type Recipe struct {
	Model
	User          User   `json:"-" gorm:"association_foreignkey:UserID"`
	UserID        uint   `json:"-" gorm:"type:int(10);unsigned;not null"`
	TitleRu       string `json:"title" gorm:"type:varchar(100);default:null"`
	TitleEn       string `json:"title" gorm:"type:varchar(100);default:null"`
	IntroRu       string `json:"intro" gorm:"type:varchar(210);default:null"`
	IntroEn       string `json:"intro" gorm:"type:varchar(210);default:null"`
	TextRu        string `json:"text" gorm:"type:text;default:null"`
	TextEn        string `json:"text" gorm:"type:text;default:null"`
	IngredientsRu string `json:"ingredients" gorm:"type:text;default:null"`
	IngredientsEn string `json:"ingredients" gorm:"type:text;default:null"`
	Slug          string `json:"slug" gorm:"type:varchar(100);not null"`
	Time          uint64 `json:"time" gorm:"type:smallint(5);unsigned;default:'0'"`
	Image         string `json:"image" gorm:"type:varchar(200);unsigned;default:'default.jpg'"`
	ReadyRu       byte   `json:"ready" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	ReadyEn       byte   `json:"ready" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	ApprovedRu    byte   `json:"approved" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	ApprovedEn    byte   `json:"approved" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	PublishedRu   byte   `json:"published" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	PublishedEn   byte   `json:"published" gorm:"type:tinyint(1);unsigned;default:'0';not null"`
	Simple        byte   `json:"simple" form:"type:tinyint(1);unsigned;default:'0';not null"`
}
