package store_test

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	t.Parallel()

	st, teardown := store.TestStore(t, databaseURL)
	defer teardown("users", "recipes")

	user, err := st.User().Create(&model.User{
		Email:    "user@example.com",
		Password: "111",
	})

	if err != nil {
		t.Error(err)
	}

	if user == nil {
		t.Error("Method Create must return user but nil returned")
	}
}
