package model

import "testing"

// TestUser helper function for testing
func TestUser(t *testing.T) *User {
	return &User{
		Name:     "Anna",
		Email:    "anna@mail.com",
		Password: "password",
	}
}
