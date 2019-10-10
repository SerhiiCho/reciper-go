package model_test

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	user := model.TestUser(t)

	if user.BeforeCreate() != nil {
		t.Error("No errors must be returned")
	}

	if user.HashedPassword == "" {
		t.Error("Hashed password must not be empty")
	}
}
