package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/SerhiiCho/reciper/backend/utils"

	"github.com/SerhiiCho/reciper/backend/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

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
	serv.router.Use(serv.appMiddleware, serv.setRequestID, serv.logRequest)
	serv.router.HandleFunc("/api/sessions", serv.sessionCreate()).Methods("POST")
	serv.router.HandleFunc("/api/recipes", serv.recipeIndex()).Methods("GET")
	serv.router.HandleFunc("/api/users", serv.userCreate()).Methods("POST")

	auth := serv.router.NewRoute().Subrouter()
	auth.Use(serv.authenticateUser)
	auth.HandleFunc("/api/recipes", serv.recipeCreate()).Methods("POST")
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
		utils.HandleError("json encode error", err, "")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
