package apiserver

import (
	"encoding/json"
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/SerhiiCho/reciper/backend/utils"
	"net/http"
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

func (serv *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	serv.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (serv *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		utils.HandleError("json encode error in user_create@respond", err, "")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
