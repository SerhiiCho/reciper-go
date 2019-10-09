package api

import (
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/app/middleware"
	"github.com/gin-gonic/gin"
)

// API struct
type API struct {
	App    *app.App
	config *Config
	router *gin.Engine
}

// NewAPI creates app struct
func NewAPI(app *app.App, config *Config) *API {
	return &API{
		App:    app,
		config: config,
		router: gin.Default(),
	}
}

// Start configures routes and start the server
func (api *API) Start() error {
	api.configureRouter()
	api.router.Use(middleware.GetAll()...)

	return api.router.Run(api.config.BindAddr)
}

// configureRouter adds routes
func (api *API) configureRouter() {
	api.router.GET("/api/recipes", api.RecipeIndex())
	api.router.POST("/api/recipes", api.RecipeCreate())
}
