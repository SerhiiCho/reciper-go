package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RecipeIndex handles GET request on showing the list of all recipes
func (api *API) RecipeIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, api.App.Database.GetRecipes())
	}
}
