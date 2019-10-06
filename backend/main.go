package main

import (
	"database/sql"
	"github.com/SerhiiCho/reciper/backend/http/handler"
	"github.com/SerhiiCho/reciper/backend/http/middleware"
	"github.com/SerhiiCho/reciper/backend/models"
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
	recipeRepo := models.NewRecipe()
	db, err := sql.Open("mysql", "root@tcp(reciper_db:1001)/reciper?charset=utf8")
	utils.HandleError("Database connection error", err)

	r.Use(middleware.App())

	r.GET("/api/recipes", handler.RecipesIndex(recipeRepo, db))
	r.POST("/api/recipes", handler.RecipesCreate(recipeRepo, db))

	utils.HandleError("Can't serve the app", r.Run())
}
