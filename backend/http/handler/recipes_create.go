package handler

import (
	"github.com/SerhiiCho/reciper/backend/models"
	"github.com/SerhiiCho/reciper/backend/repos"
	"net/http"
	"strconv"

	"github.com/SerhiiCho/reciper/backend/utils"

	"github.com/gin-gonic/gin"
)

// RecipesCreate handles POST request on creating a new recipe item
func RecipesCreate(recipeRepo *repos.RecipeRepo) gin.HandlerFunc {

	return func(c *gin.Context) {
		time, parseErr := strconv.ParseUint(c.PostForm("time"), 10, 32)
		utils.HandleError("Error parsing time from request", parseErr)

		recipeRepo.Add(models.Recipe{
			Title:       c.PostForm("title"),
			Excerpt:     utils.StrLimit(c.PostForm("title"), 42),
			Intro:       c.PostForm("intro"),
			Text:        c.PostForm("text"),
			Ingredients: c.PostForm("ingredients"),
			Slug:        c.PostForm("slug"),
			Time:        time,
			Image:       c.PostForm("image"),
		})

		c.Status(http.StatusNoContent)
	}
}
