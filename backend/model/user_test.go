package model_test

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"testing"
)

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

func TestUser_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{"valid", func() *model.User {
			return model.TestUser(t)
		}, true},
		{"email invalid 1", func() *model.User {
			user := model.TestUser(t)
			user.Email = "email@mailcom"
			return user
		}, false},
		{"email invalid 2", func() *model.User {
			user := model.TestUser(t)
			user.Email = "anna.mail.com"
			return user
		}, false},
		{"email empty", func() *model.User {
			user := model.TestUser(t)
			user.Email = ""
			return user
		}, false},
		{"name empty", func() *model.User {
			user := model.TestUser(t)
			user.Name = ""
			return user
		}, true},
		{"name too short", func() *model.User {
			user := model.TestUser(t)
			user.Name = "na"
			return user
		}, false},
		{"name too long", func() *model.User {
			user := model.TestUser(t)
			user.Name = "some very long string with more than 50 charactersX" // 51
			return user
		}, false},
		{"name short okey", func() *model.User {
			user := model.TestUser(t)
			user.Name = "nag"
			return user
		}, true},
		{"password empty", func() *model.User {
			user := model.TestUser(t)
			user.Password = ""
			return user
		}, false},
		{"password short okey", func() *model.User {
			user := model.TestUser(t)
			user.Password = "12345678"
			return user
		}, true},
		{"password long okey", func() *model.User {
			user := model.TestUser(t)
			user.Password = "usdfjdsfljdsflksaj fjdsklfjdslkf jdkfjsdlfk;ajdfkl dklf asd;fdlskflaskdfsdfasdf dsfsadfdsaf dasfasdfdusdfjdsfljdsflksaj fjdsklfjdslkf jdkfjsdlfk;ajdfkl dklf asd;fdlskflaskdfsdfasdf dsfsadfdsaf dasfasdfddasfasdfddasfasdfddasfasdfddasfasdfddasfasdfddas"
			return user
		}, true},
		{"password too short", func() *model.User {
			user := model.TestUser(t)
			user.Password = "1234567"
			return user
		}, false},
		{"password too long", func() *model.User {
			user := model.TestUser(t)
			user.Password = "usdfjdsfljdsflksaj fjdsklfjdslkf jdkfjsdlfk;ajdfkl dklf asd;fdlskflaskdfsdfasdf dsfsadfdsaf dasfasdfdusdfjdsfljdsflksaj fjdsklfjdslkf jdkfjsdlfk;ajdfkl dklf asd;fdlskflaskdfsdfasdf dsfsadfdsaf dasfasdfddasfasdfddasfasdfddasfasdfddasfasdfddasfasdfddasX"
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
