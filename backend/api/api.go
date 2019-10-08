package api

import (
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/gin-gonic/gin"
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

	utils.HandleError("Can't serve the app", router.Run(), "")
}
