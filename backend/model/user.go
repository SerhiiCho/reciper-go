package model

import (
	"github.com/SerhiiCho/reciper/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Model
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	Email          string  `json:"email"`
	Username       string  `json:"username"`
	XP             uint    `json:"xp"`
	StreakDays     uint    `json:"streak_days"`
	Popularity     float64 `json:"popularity"`
	Active         bool    `json:"active"`
	StreakCheck    string  `json:"streak_check"`
	NotifCheck     string  `json:"notif_check"`
	OnlineCheck    string  `json:"online_check"`
	ContactCheck   string  `json:"contact_check"`
	HashedPassword []byte
	Photo          string `json:"photo"`
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
