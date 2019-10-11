package apiserver

import (
	"errors"
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

func (serv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv.router.ServeHTTP(w, r)
}

func (serv server) configureRouter() {
	serv.router.Use(serv.appMiddleware)
	serv.router.HandleFunc("/api/sessions", serv.sessionCreate()).Methods("POST")
	serv.router.HandleFunc("/api/recipes", serv.recipeIndex()).Methods("GET")
	serv.router.HandleFunc("/api/users", serv.userCreate()).Methods("POST")

	auth := serv.router.NewRoute().Subrouter()
	auth.Use(serv.authenticateUser)
	auth.HandleFunc("/api/recipes", serv.recipeCreate()).Methods("POST")
}
