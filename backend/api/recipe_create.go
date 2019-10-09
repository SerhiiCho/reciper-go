package api

import (
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"net/http"
	"os"
	"strconv"
	"time"

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
			Image:         uploadImage(ctx),
		})

		ctx.Status(http.StatusNoContent)
	}
}

func uploadImage(ctx *gin.Context) string {
	fileName := strconv.Itoa(int(time.Now().Unix())) + ".jpg"

	imageFile, imageFileErr := ctx.FormFile("image")
	utils.HandleError("Form file error", imageFileErr, "")

	file, fileOpenErr := imageFile.Open()
	utils.HandleError("File open error", fileOpenErr, "")

	decodedImage, _, decodeErr := image.Decode(file)
	utils.HandleError("File decode error", decodeErr, "")

	smImage := resize.Resize(900, 600, decodedImage, resize.Lanczos3)
	lgImage := resize.Resize(225, 150, decodedImage, resize.Lanczos3)

	smImageOut, smErr := os.Create("storage/recipes/small/" + fileName)
	lgImageOut, lgErr := os.Create("storage/recipes/large/" + fileName)

	utils.HandleError("Error creating small recipe image", smErr, "")
	utils.HandleError("Error creating large recipe image", lgErr, "")

	smImageErr := jpeg.Encode(smImageOut, smImage, nil)
	lgImageErr := jpeg.Encode(lgImageOut, lgImage, nil)

	utils.HandleError("Error encoding small recipe image", smImageErr, "")
	utils.HandleError("Error encoding large recipe image", lgImageErr, "")

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
