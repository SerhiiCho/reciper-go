package apiserver

import (
	"context"
	"errors"
	"github.com/SerhiiCho/reciper/backend/apiserver/middleware"
	"github.com/SerhiiCho/reciper/backend/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

const (
	sessionName               = "reciper"
	contextKeyUser contextKey = iota
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type contextKey int8

type server struct {
	router       *mux.Router
	store        store.Store
	sessionStore sessions.Store
}

// newServer configures router and returns pointer to server struct
func newServer(store store.Store, sessionStore sessions.Store) *server {
	serv := &server{
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
	}

	serv.configureRouter()

	return serv
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

func (serv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv.router.ServeHTTP(w, r)
}

func (serv server) configureRouter() {
	serv.router.Use(middleware.AppMiddleware)
	serv.router.HandleFunc("/api/sessions", serv.sessionCreate()).Methods("POST")
	serv.router.HandleFunc("/api/recipes", serv.recipeIndex()).Methods("GET")
	serv.router.HandleFunc("/api/recipes", serv.recipeCreate()).Methods("POST")
	serv.router.HandleFunc("/api/users", serv.userCreate()).Methods("POST")
}
