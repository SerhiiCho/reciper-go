package api

import (
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/gin-gonic/gin"
)

type API struct {
	App *app.App
}

func NewAPI(app *app.App) *API {
	return &API{App: app}
}

func (a *API) Init(router *gin.Engine) {
	router.GET("/recipes", RecipeIndex())
	router.POST("/recipes", RecipeCreate())

	utils.HandleError("Can't serve the app", router.Run(), "")
}
