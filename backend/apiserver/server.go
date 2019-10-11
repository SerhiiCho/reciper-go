package apiserver

import (
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

//
func (serv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serv.router.ServeHTTP(w, r)
}

func (serv server) configureRouter() {

}
