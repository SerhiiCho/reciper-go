package apiserver

import (
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/app/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// APIServer struct
type APIServer struct {
	App    *app.App
	config *Config
	router *mux.Router
}

// NewAPIServer creates app struct
func NewAPIServer(app *app.App, config *Config) *APIServer {
	return &APIServer{
		App:    app,
		config: config,
		router: mux.NewRouter().PathPrefix("/api").Subrouter(),
	}
}

// Start configures routes and start the server
func (api *APIServer) Start() error {
	api.configureRouter()
	api.router.Use(middleware.AppMiddleware)

	log.Printf("Server started at %s", api.config.BindAddr)

	return http.ListenAndServe(api.config.BindAddr, api.router)
}

// configureRouter adds routes
func (api *APIServer) configureRouter() {
	api.router.HandleFunc("/recipes", api.recipeIndex()).Methods("GET")
	api.router.HandleFunc("/recipes", api.recipeCreate()).Methods("POST")
}
