package main

import (
	"database/sql"
	"fmt"
	"github.com/SerhiiCho/reciper/backend/http/handler"
	"github.com/SerhiiCho/reciper/backend/http/middleware"
	"github.com/SerhiiCho/reciper/backend/models"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	utils.HandleError("Error loading .env file", err)
}

func main() {
	r := gin.Default()
	recipeRepo := models.NewRecipe()
	db := getDB()

	r.Use(middleware.App())

	r.GET("/api/recipes", handler.RecipesIndex(recipeRepo, db))
	r.POST("/api/recipes", handler.RecipesCreate(recipeRepo, db))

	utils.HandleError("Can't serve the app", r.Run())
}

func getDB() *sql.DB {
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, host, port, name)

	db, err := sql.Open("mysql", dataSource)
	utils.HandleError("Database connection error", err)

	return db
}
