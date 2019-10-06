package handler

import (
	"github.com/SerhiiCho/reciper/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecipesIndex handles GET request on showing the list of all recipes
func RecipesIndex(recipeRepo *models.RecipeRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, recipeRepo.IndexRecipe())
	}
}
