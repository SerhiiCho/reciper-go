package apiserver

import (
	"net/http"

	"github.com/SerhiiCho/reciper/backend/model"
)

// userCreate adds new user to the store
func (serv *server) userCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		if err := serv.store.User().CreateUser(user); err != nil {
			serv.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		user.SanitizeUserFields()
		serv.respond(w, r, http.StatusCreated, user)
	}
}
