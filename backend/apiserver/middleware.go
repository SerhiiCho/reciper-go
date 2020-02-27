package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

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

func (serv *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-Id", id)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextKeyRequestID, id)))
	})
}

func (serv *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		respWriter := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(respWriter, r)

		fmt.Printf(
			"[INFO]: %s %s [%d %s] in %v | %s | %s | %s\n",
			r.Method,
			r.RequestURI,
			respWriter.code,
			http.StatusText(respWriter.code),
			time.Now().Sub(start),
			r.RemoteAddr,
			time.Now().Format("02.01.2006 15:04:05"),
			r.Context().Value(contextKeyRequestID),
		)
	})
}
