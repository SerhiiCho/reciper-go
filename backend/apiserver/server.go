package apiserver

import (
	"github.com/SerhiiCho/reciper/backend/apiserver/middleware"
	"github.com/SerhiiCho/reciper/backend/store"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	router *mux.Router
	store  store.Store
}

// newServer configures router and returns pointer to server struct
func newServer(store store.Store) *server {
	serv := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	serv.configureRouter()

	return serv
}

func (serv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv.router.ServeHTTP(w, r)
}

func (serv server) configureRouter() {
	serv.router.Use(middleware.AppMiddleware)
	serv.router.HandleFunc("/api/recipes", serv.recipeIndex()).Methods("GET")
	serv.router.HandleFunc("/api/recipes", serv.recipeCreate()).Methods("POST")
	serv.router.HandleFunc("/api/users", serv.userCreate()).Methods("POST")
}
