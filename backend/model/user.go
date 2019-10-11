package model

import (
	valid "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Model
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	Email          string  `json:"email"`
	XP             uint    `json:"xp"`
	StreakDays     uint    `json:"streak_days"`
	Popularity     float64 `json:"popularity"`
	Active         bool    `json:"active"`
	StreakCheck    string  `json:"streak_check"`
	NotifCheck     string  `json:"notif_check"`
	OnlineCheck    string  `json:"online_check"`
	ContactCheck   string  `json:"contact_check"`
	Photo          string  `json:"photo"`
	Password       string  `json:"password,omitempty"`
	HashedPassword string  `json:"-"`
}

// Validate validates user data
func (user *User) Validate() error {
	return valid.ValidateStruct(
		user,
		valid.Field(&user.Name, valid.Length(3, 50)),
		valid.Field(
			&user.Email,
			valid.Required.Error("Эл. адрес обязателен к заполнению"),
			valid.Length(7, 190).Error("Длина не пойдет"),
			is.Email.Error("Не валидная почта"),
		),
		valid.Field(&user.Password, valid.By(requiredIf(user.HashedPassword == "")), valid.Length(8, 250)),
	)
}

// SanitizeUserFields sets some user fields to empty string
func (user *User) SanitizeUserFields() {
	user.Password = ""
}

// BeforeCreate executes before new user is created
func (user *User) BeforeCreate() error {
	if len(user.Password) > 0 {
		newPwd, err := generatePasswordHash(user.Password)

		if err != nil {
			return err
		}

		user.HashedPassword = newPwd
	}

	return nil
}

// generatePasswordHash generates encrypted user password
func generatePasswordHash(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hashedPwd), nil
}
