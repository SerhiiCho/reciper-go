package handler

import (
	"github.com/SerhiiCho/reciper/backend/models/recipe"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecipesGET handles GET request on showing the list of all recipes
func RecipesGET(recipe *recipe.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "604800")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Access-Control-Request-Method")
		c.Header("Access-Control-Allow-Methods", "GET, POST")

		c.JSON(http.StatusOK, recipe.GetAll())
	}
}
