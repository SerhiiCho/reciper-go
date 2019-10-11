package apiserver

import (
	"context"
	"net/http"
)

// appMiddleware sets http headers
func (serv server) appMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Max-Age", "604800")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Access-Control-Request-Method")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

// authenticateUser middleware
func (serv server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, sessionErr := serv.sessionStore.Get(r, sessionName)

		if sessionErr != nil {
			serv.error(w, r, http.StatusInternalServerError, sessionErr)
			return
		}

		id, ok := session.Values["user_id"]

		if !ok {
			serv.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		user, err := serv.store.User().FindUser(id.(uint))

		if err != nil {
			serv.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextKeyUser, user)))
	})
}
