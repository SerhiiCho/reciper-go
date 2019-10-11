package apiserver

import (
	"errors"
	"net/http"
)

var errIncorrectEmailOrPassword = errors.New("incorrect email or password")
var sessionName = "reciper"

func (serv *server) sessionCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := serv.store.User().FindByEmail(email)

		if err != nil || !user.ComparePassword(password) {
			serv.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, sessionErr := serv.sessionStore.Get(r, sessionName)

		if sessionErr != nil {
			serv.error(w, r, http.StatusInternalServerError, sessionErr)
		}

		session.Values["user_id"] = user.ID

		if err := serv.sessionStore.Save(r, w, session); err != nil {
			serv.error(w, r, http.StatusInternalServerError, sessionErr)
		}

		serv.respond(w, r, http.StatusOK, nil)
	}
}
