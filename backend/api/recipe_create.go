package api

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/SerhiiCho/reciper/backend/utils"

	"github.com/gin-gonic/gin"
)

// RecipeCreate handles POST request on creating a new recipe item
func (api *API) RecipeCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, fileErr := ctx.FormFile("image")
		utils.HandleError("Form file error", fileErr, "")

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
			Image:         uploadImage(file, ctx),
		})

		ctx.Status(http.StatusNoContent)
	}
}

func uploadImage(file *multipart.FileHeader, ctx *gin.Context) string {
	fileName := strconv.Itoa(int(time.Now().Unix())) + ".jpg"

	uploadErr := ctx.SaveUploadedFile(file, "storage/"+fileName)
	utils.HandleError("File uploading error", uploadErr, "")

	return fileName
}

func setTimeField(value string) uint64 {
	cookTime, parseErr := strconv.ParseUint(value, 10, 32)
	utils.HandleError("Error parsing time from request", parseErr, "")
	return cookTime
}

func setReadyField(value string) byte {
	if value == "1" {
		return byte(1)
	}
	return byte(0)
}
