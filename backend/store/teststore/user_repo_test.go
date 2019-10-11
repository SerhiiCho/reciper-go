package teststore_test

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store/teststore"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	st := teststore.New()
	userErr := st.User().Create(model.TestUser(t))

	if userErr != nil {
		t.Error("Method Create must return nil but error returned")
	}
}

func TestUserRepo_FindByEmail(t *testing.T) {
	st := teststore.New()
	email := "anna@mail.com"

	t.Run("user doesn't exist in db", func(t *testing.T) {
		user, err1 := st.User().FindByEmail(email)

		if err1 == nil {
			t.Error("FindByEmail must return error because user doesn't exist")
		}

		if user != nil {
			t.Error("returned user must be nil")
		}
	})

	t.Run("user exist in db", func(t *testing.T) {
		errUserCreate := st.User().Create(model.TestUser(t))

		if errUserCreate != nil {
			t.Fatal("Can't create user in database", errUserCreate)
		}

		user, err2 := st.User().FindByEmail(email)

		if err2 != nil {
			t.Error("FindByEmail should not return error because user exist in db")
		}

		if user == nil || user.Email != email {
			t.Errorf("user with email %s must be returned", email)
		}
	})
}
