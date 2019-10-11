package sqlstore_test

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store"
	"github.com/SerhiiCho/reciper/backend/store/sqlstore"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users", "recipes")

	st := sqlstore.New(db)
	userErr := st.User().CreateUser(model.TestUser(t))

	if userErr != nil {
		t.Error("Method CreateUser must return nil but error returned")
	}
}

func TestUserRepo_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users", "recipes")

	st := sqlstore.New(db)
	email := "anna@mail.com"

	t.Run("user doesn't exist in db", func(t *testing.T) {
		user, findErr := st.User().FindByEmail(email)

		if findErr != store.ErrRecordNotFound {
			t.Error("FindByEmail must return ErrRecordNotFound error because user doesn't exist")
		}

		if user != nil {
			t.Error("returned user must be nil")
		}
	})

	t.Run("user exist in db", func(t *testing.T) {
		errUserCreate := st.User().CreateUser(model.TestUser(t))

		if errUserCreate != nil {
			t.Fatal("Can't create user in database", errUserCreate)
		}

		user, notFoundErr := st.User().FindByEmail(email)

		if notFoundErr != nil {
			t.Error("FindByEmail must return nil instead of error because user exist in database")
		}

		if user == nil {
			t.Errorf("user with email %s must be returned instead of nil", email)
		}

		if user != nil && user.Email != email {
			t.Errorf("user email must be equal to %s but instead of %s", email, user.Email)
		}

	})
}
