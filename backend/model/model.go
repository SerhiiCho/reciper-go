package model

type Model struct {
	ID        uint   `json:"id" gorm:"type:int(10);unsigned;not null;AUTO_INCREMENT;primary_key" `
	CreatedAt string `json:"created_at" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt string `json:"updated_at" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	DeletedAt string `json:"updated_at" gorm:"type:datetime;default:null"`
}
