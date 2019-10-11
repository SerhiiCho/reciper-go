package teststore_test

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/store"
	"github.com/SerhiiCho/reciper/backend/store/teststore"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	st := teststore.New()
	userErr := st.User().CreateUser(model.TestUser(t))

	if userErr != nil {
		t.Error("Method CreateUser must return nil but error returned")
	}
}

func TestUserRepo_FindByEmail(t *testing.T) {
	st := teststore.New()
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

func TestUserRepo_FindUser(t *testing.T) {
	st := teststore.New()

	t.Run("user doesn't exist in db", func(t *testing.T) {
		user, findErr := st.User().FindUser(32)

		if findErr != store.ErrRecordNotFound {
			t.Error("FindByEmail must return ErrRecordNotFound error because user doesn't exist")
		}

		if user != nil {
			t.Error("returned user must be nil")
		}
	})

	t.Run("user exist in db", func(t *testing.T) {
		testUser := model.TestUser(t)
		errUserCreate := st.User().CreateUser(testUser)

		if errUserCreate != nil {
			t.Fatal("Can't create user in database", errUserCreate)
		}

		user, notFoundErr := st.User().FindUser(testUser.ID)

		if notFoundErr != nil {
			t.Error("FindByEmail must return nil instead of error because user exist in database")
		}

		if user == nil {
			t.Errorf("user with id %d must be returned instead of nil", testUser.ID)
		}

		if user != nil && user.ID != testUser.ID {
			t.Errorf("user id must be equal to %d but instead of %d", testUser.ID, user.ID)
		}

	})
}
