package model

import "testing"

// TestUser helper function for testing
func TestUser(t *testing.T) *User {
	return &User{
		Name:     "Anna",
		Email:    "anna_korotchaeva@mail.com",
		Password: "password",
	}
}
