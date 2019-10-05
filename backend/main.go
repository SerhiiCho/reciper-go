package main

import (
	"github.com/SerhiiCho/reciper/backend/httpd/handler"
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

	r.GET("/api/recipes", handler.RecipesGET())

	err := r.Run()
	utils.HandleError("Can't serve the app", err)
}
