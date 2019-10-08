package model

import (
	"github.com/SerhiiCho/reciper/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Model
	Name           string  `json:"name" gorm:"type:varchar(50);default:null"`
	Status         string  `json:"status" gorm:"type:varchar(250);default:null"`
	Email          string  `json:"email" gorm:"type:varchar(250);default:null"`
	Username       string  `json:"username" gorm:"type:varchar(40);unique;not null"`
	XP             uint    `json:"xp" gorm:"type:smallint;default:'0'"`
	StreakDays     uint    `json:"streak_days" gorm:"type:int(8);default:'0'"`
	Popularity     float64 `json:"popularity" gorm:"type:decimal(8,1);default:'0'"`
	Active         bool    `json:"active" gorm:"type:tinyint(1);default:'1'"`
	StreakCheck    string  `json:"streak_check" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	NotifCheck     string  `json:"notif_check" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	OnlineCheck    string  `json:"online_check" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ContactCheck   string  `json:"contact_check" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	HashedPassword []byte
	Photo          string `json:"photo" gorm:"type:varchar(200);unsigned;default:'default.jpg'"`
}

// GeneratePasswordHash function
func GeneratePasswordHash(password []byte) []byte {
	password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	utils.HandleError("Error after generating from password using bcrypt package", err, "")

	return password
}

// ComparePasswordHash function
func ComparePasswordHash(hashedPassword, givenPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, givenPassword)
	return err == nil
}

// SetPassword sets password
func (user *User) SetPassword(password string) {
	hashed := GeneratePasswordHash([]byte(password))
	user.HashedPassword = hashed
}

// CheckPassword checks password and returns true if password match
func (user *User) CheckPassword(password string) bool {
	return ComparePasswordHash(user.HashedPassword, []byte(password))
}
