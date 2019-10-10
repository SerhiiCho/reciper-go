package store_test

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	st, teardown := store.TestStore(t, databaseURL)
	defer teardown("users", "recipes")

	email := "user@example.com"

	user, err := st.User().Create(&model.User{Email: email})

	if err != nil {
		t.Error(err)
	}

	if user == nil || user.Email != email {
		t.Errorf("Method Create must return user with email %s", email)
	}
}

func TestUserRepo_FindByEmail(t *testing.T) {
	st, teardown := store.TestStore(t, databaseURL)
	defer teardown("users", "recipes")

	email := "user@example.com"

	t.Run("user doesn't exist in db", func(t *testing.T) {
		err1 := st.User().FindByEmail(email)

		if err1 == nil {
			t.Error("FindByEmail must return error because user doesn't exist")
		}
	})

	t.Run("user exist in db", func(t *testing.T) {
		user := &model.User{Email: email}
		_, errUserCreate := st.User().Create(user)

		if errUserCreate != nil {
			t.Fatal("Can't create user in database", errUserCreate)
		}

		err2 := st.User().FindByEmail(email)

		if err2 != nil {
			t.Error("FindByEmail should not return error because user exist in db")
		}

		if user.Email != email {
			t.Errorf("user email must be %s", email)
		}
	})
}
