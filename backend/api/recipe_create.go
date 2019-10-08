package api

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"net/http"
	"strconv"

	"github.com/SerhiiCho/reciper/backend/utils"

	"github.com/gin-gonic/gin"
)

// RecipeCreate handles POST request on creating a new recipe item
func (api *API) RecipeCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		time, parseErr := strconv.ParseUint(ctx.PostForm("time"), 10, 32)
		utils.HandleError("Error parsing time from request", parseErr, "")

		api.App.Database.NewRecord(model.Recipe{
			TitleRu:       ctx.PostForm("title"),
			TitleEn:       ctx.PostForm("title"),
			IntroRu:       ctx.PostForm("intro"),
			IntroEn:       ctx.PostForm("intro"),
			TextRu:        ctx.PostForm("text"),
			TextEn:        ctx.PostForm("text"),
			IngredientsRu: ctx.PostForm("ingredients"),
			IngredientsEn: ctx.PostForm("ingredients"),
			Slug:          ctx.PostForm("slug"),
			Time:          time,
			Image:         ctx.PostForm("image"),
		})

		ctx.Status(http.StatusNoContent)
	}
}
