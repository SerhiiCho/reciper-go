package main

import (
	"github.com/SerhiiCho/reciper/backend/http/handler"
	"github.com/SerhiiCho/reciper/backend/http/middleware"
	"github.com/SerhiiCho/reciper/backend/repos"
	"github.com/SerhiiCho/reciper/backend/storage"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError("Error loading .env file", err, "")
}

func main() {
	router := gin.Default()
	DB := storage.GetDB()
	recipeRepo := repos.NewRecipeRepo(DB)

	router.Use(middleware.AppMiddle())

	router.GET("/api/recipes", handler.RecipesIndex(recipeRepo))
	router.POST("/api/recipes", handler.RecipesCreate(recipeRepo))

	utils.HandleError("Can't serve the app", router.Run(), "")
}
