package model_test

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{"valid", func() *model.User {
			return model.TestUser(t)
		}, true},
		{"empty email", func() *model.User {
			user := model.TestUser(t)
			user.Email = ""
			return user
		}, false},
		{"name can be empty", func() *model.User {
			user := model.TestUser(t)
			user.Name = ""
			return user
		}, true},
		{"empty password", func() *model.User {
			user := model.TestUser(t)
			user.Password = ""
			return user
		}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				// if validate returns error
				if tc.user().Validate() != nil {
					t.Error("Validate must return nil but error returned")
				}
			} else {
				// if validate returns nil
				if tc.user().Validate() == nil {
					t.Error("Validate must return error but nil returned")
				}
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	t.Parallel()

	user := model.TestUser(t)

	if user.BeforeCreate() != nil {
		t.Error("No errors must be returned")
	}

	if user.HashedPassword == "" {
		t.Error("Hashed password must not be empty")
	}
}
