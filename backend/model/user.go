package model

import "golang.org/x/crypto/bcrypt"

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

func GeneratePasswordHash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func ComparePasswordHash(hashedPassword, givenPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, givenPassword)
	return err == nil
}

func (user *User) SetPassword(password string) error {
	hashed, err := GeneratePasswordHash([]byte(password))
	if err != nil {
		return err
	}
	user.HashedPassword = hashed
	return nil
}

func (user *User) CheckPassword(password string) bool {
	return ComparePasswordHash(user.HashedPassword, []byte(password))
}
