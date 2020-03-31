package apiserver

import (
	"fmt"
	"image"
	"image/jpeg"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/SerhiiCho/reciper/backend/model"
	"github.com/nfnt/resize"

	"github.com/SerhiiCho/reciper/backend/utils"
)

// recipeCreate handles POST request on creating a new recipe item
func (serv *server) recipeCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serv.respond(w, r, http.StatusOK, r.Context().Value(contextKeyUser).(*model.User))
		//serv.store.CreateRecipe(&model.Recipe{
		//	TitleRu:       r.FormValue("title-ru"),
		//	TitleEn:       r.FormValue("title-en"),
		//	IntroRu:       r.FormValue("intro-ru"),
		//	IntroEn:       r.FormValue("intro-en"),
		//	TextRu:        r.FormValue("text-ru"),
		//	TextEn:        r.FormValue("text-en"),
		//	IngredientsRu: r.FormValue("ingredients-ru"),
		//	IngredientsEn: r.FormValue("ingredients-en"),
		//	Slug:          r.FormValue("slug"),
		//	Time:          setTimeField(r.FormValue("time")),
		//	ReadyRu:       setReadyField(r.FormValue("ready-ru")),
		//	ReadyEn:       setReadyField(r.FormValue("ready-en")),
		//	Image:         uploadImage(r),
		//})

		//w.WriteHeader(http.StatusNoContent)
	}
}

// uploadImage uploads images to storage
func uploadImage(r *http.Request) string {
	_, imageFile, imageFileErr := r.FormFile("image")
	utils.HandleError("Form file error", imageFileErr, "")

	file, fileOpenErr := imageFile.Open()
	utils.HandleError("File open error", fileOpenErr, "")

	decodedImage, _, decodeErr := image.Decode(file)
	utils.HandleError("File decode error", decodeErr, "")

	smImage := resize.Resize(225, 150, decodedImage, resize.Lanczos3)
	lgImage := resize.Resize(900, 600, decodedImage, resize.Lanczos3)

	filePath, dirPath := generateFileNameAndFilePath()

	createNeededDirectories(dirPath)
	createFiles(filePath, smImage, lgImage)

	return filePath
}

// createFiles creates files in storage
func createFiles(filePath string, smImage image.Image, lgImage image.Image) {
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
}

// createNeededDirectories creates directories in storage
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
