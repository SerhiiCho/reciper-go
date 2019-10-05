package handler

import (
	"github.com/SerhiiCho/reciper/backend/models"
	"net/http"
	"strconv"

	"github.com/SerhiiCho/reciper/backend/utils"

	"github.com/gin-gonic/gin"
)

// RecipesCreate handles POST request on creating a new recipe item
func RecipesCreate(recipeRepo *models.RecipeRepo) gin.HandlerFunc {

	return func(c *gin.Context) {
		time, parseErr := strconv.ParseUint(c.PostForm("time"), 10, 32)
		utils.HandleError("Error parsing time from request", parseErr)

		recipeRepo.CreateRecipe(models.Recipe{
			Title:       c.PostForm("title"),
			Intro:       c.PostForm("intro"),
			Text:        c.PostForm("text"),
			Ingredients: c.PostForm("ingredients"),
			Slug:        c.PostForm("slug"),
			Time:        time,
			Image:       c.PostForm("image"),
		})

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "604800")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Access-Control-Request-Method")
		c.Header("Access-Control-Allow-Methods", "GET, POST")

		c.Status(http.StatusNoContent)
	}
}
