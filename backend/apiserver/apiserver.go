package apiserver

import (
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/app/middleware"
	"github.com/gorilla/mux"
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
		router: mux.NewRouter(),
	}
}

// Start configures routes and start the server
func (api *APIServer) Start() error {
	api.configureRouter()
	api.router.Use(middleware.AppMiddleware)

	return http.ListenAndServe(api.config.BindAddr, api.router)
}

// configureRouter adds routes
func (api *APIServer) configureRouter() {
	api.router.HandleFunc("/api/recipes", api.recipeIndex()).Methods("GET")
	api.router.HandleFunc("/api/recipes", api.recipeCreate()).Methods("POST")
}
