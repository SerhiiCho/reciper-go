package api

import (
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/gin-gonic/gin"
	"log"
)

// API struct
type API struct {
	App *app.App
}

// NewAPI creates app struct
func NewAPI(app *app.App) *API {
	return &API{App: app}
}

// Init initializes the api routes
func (api *API) Init(router *gin.Engine) {
	router.GET("/api/recipes", api.RecipeIndex())
	router.POST("/api/recipes", api.RecipeCreate())

	if err := router.Run(); err != nil {
		log.Fatal("Can't serve the app")
	}
}
