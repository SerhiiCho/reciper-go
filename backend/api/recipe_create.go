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
		api.App.Database.CreateRecipe(&model.Recipe{
			TitleRu:       ctx.PostForm("title-ru"),
			TitleEn:       ctx.PostForm("title-en"),
			IntroRu:       ctx.PostForm("intro-ru"),
			IntroEn:       ctx.PostForm("intro-en"),
			TextRu:        ctx.PostForm("text-ru"),
			TextEn:        ctx.PostForm("text-en"),
			IngredientsRu: ctx.PostForm("ingredients-ru"),
			IngredientsEn: ctx.PostForm("ingredients-en"),
			Slug:          ctx.PostForm("slug"),
			Time:          setTimeField(ctx.PostForm("time")),
			ReadyRu:       setReadyField(ctx.PostForm("ready-ru")),
			ReadyEn:       setReadyField(ctx.PostForm("ready-en")),
			Image:         ctx.PostForm("image"),
		})

		ctx.Status(http.StatusNoContent)
	}
}

func setTimeField(value string) uint64 {
	time, parseErr := strconv.ParseUint(value, 10, 32)
	utils.HandleError("Error parsing time from request", parseErr, "")
	return time
}

func setReadyField(value string) byte {
	if value == "1" {
		return byte(1)
	}
	return byte(0)
}
