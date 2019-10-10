package apiserver

import (
	"github.com/SerhiiCho/reciper/backend/apiserver/middleware"
	"github.com/SerhiiCho/reciper/backend/store"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// APIServer struct
type APIServer struct {
	config *Config
	router *mux.Router
	store  *store.Store
}

// NewAPIServer returns pointer on APIServer
func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter().PathPrefix("/api").Subrouter(),
	}
}

// Start configures routes and start the server
func (api *APIServer) Start() error {
	api.configureRouter()
	api.configureStore()

	log.Printf("Server started at %s", api.config.BindAddr)

	return http.ListenAndServe(api.config.BindAddr, api.router)
}

// configureRouter adds routes
func (api *APIServer) configureRouter() {
	api.router.Use(middleware.AppMiddleware)
	api.router.HandleFunc("/recipes", api.recipeIndex()).Methods("GET")
	api.router.HandleFunc("/recipes", api.recipeCreate()).Methods("POST")
}

// configureStore configures database
func (api *APIServer) configureStore() {
	st := store.NewStore(api.config.Store)

	if err := st.Open(); err != nil {
		utils.FatalIfError("Store.Open() error in apiserver@configureStore", err)
	}

	api.store = st
}
