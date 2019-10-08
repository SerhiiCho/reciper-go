package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RecipeIndex handles GET request on showing the list of all recipes
func RecipeIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, recipeRepo.All())
	}
}
