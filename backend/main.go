package main

import (
	"database/sql"
	"fmt"
	"github.com/SerhiiCho/reciper/backend/http/handler"
	"github.com/SerhiiCho/reciper/backend/http/middleware"
	"github.com/SerhiiCho/reciper/backend/repos"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	utils.HandleError("Error loading .env file", err, "")
}

func main() {
	router := gin.Default()
	DB := getDB()
	recipeRepo := repos.NewRecipeRepo(DB)

	router.Use(middleware.AppMiddle())

	router.GET("/api/recipes", handler.RecipesIndex(recipeRepo))
	router.POST("/api/recipes", handler.RecipesCreate(recipeRepo))

	utils.HandleError("Can't serve the app", router.Run(), "")
}

func getDB() *sql.DB {
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, host, port, name)

	DB, err := sql.Open("mysql", dataSource)
	utils.HandleError("Database connection error", err, "")

	return DB
}
