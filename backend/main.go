package main

import (
	"github.com/SerhiiCho/reciper/backend/handler"
	"github.com/SerhiiCho/reciper/backend/models/recipe"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError("Error loading .env file", err)
}

func main() {
	r := gin.Default()
	recipeRepo := recipe.New()

	r.GET("/api/recipes", handler.RecipesGET(recipeRepo))
	r.POST("/api/recipes", handler.RecipesPOST(recipeRepo))

	utils.HandleError("Can't serve the app", r.Run())
}
