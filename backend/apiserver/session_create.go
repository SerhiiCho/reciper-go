package apiserver

import (
	"errors"
	"net/http"
)

var errIncorrectEmailOrPassword = errors.New("incorrect email or password")

func (serv *server) sessionCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := serv.store.User().FindByEmail(email)

		if err != nil || !user.ComparePassword(password) {
			serv.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		serv.respond(w, r, http.StatusOK, nil)
	}
}
