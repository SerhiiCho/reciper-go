package api

import (
	"fmt"
	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"math/rand"
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

	imageFile, imageFileErr := ctx.FormFile("image")
	utils.HandleError("Form file error", imageFileErr, "")

	file, fileOpenErr := imageFile.Open()
	utils.HandleError("File open error", fileOpenErr, "")

	decodedImage, _, decodeErr := image.Decode(file)
	utils.HandleError("File decode error", decodeErr, "")

	smImage := resize.Resize(225, 150, decodedImage, resize.Lanczos3)
	lgImage := resize.Resize(900, 600, decodedImage, resize.Lanczos3)

	filePath, dirPath := generateFileNameAndFilePath()
	createNeededDirectories(dirPath)

	filePathSm := fmt.Sprintf("storage/small/recipes/%s", filePath)
	filePathLg := fmt.Sprintf("storage/large/recipes/%s", filePath)

	smImageOut, smErr := os.Create(filePathSm)
	lgImageOut, lgErr := os.Create(filePathLg)

	utils.HandleError("Error creating small recipe image", smErr, "")
	utils.HandleError("Error creating large recipe image", lgErr, "")

	smImageErr := jpeg.Encode(smImageOut, smImage, nil)
	lgImageErr := jpeg.Encode(lgImageOut, lgImage, nil)

	utils.HandleError("Error encoding small recipe image", smImageErr, "")
	utils.HandleError("Error encoding large recipe image", lgImageErr, "")

	return filePath
}

// createNeededDirectories creates directories.
// If directories exist it will do nothing.
func createNeededDirectories(dirPath string) {
	mkdirErr1 := os.MkdirAll("storage/small/recipes/"+dirPath, os.ModePerm)
	mkdirErr2 := os.MkdirAll("storage/large/recipes/"+dirPath, os.ModePerm)

	utils.HandleError("Mkdir error", mkdirErr1, "")
	utils.HandleError("Mkdir error", mkdirErr2, "")
}

func generateFileNameAndFilePath() (string, string) {
	name := strconv.Itoa(int(time.Now().Unix()))
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	randNum := strconv.Itoa(rand.Intn(9999))

	dirPath := year + "/" + month + "/"
	filePath := dirPath + name + "-" + randNum + ".jpg"

	return filePath, dirPath
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
