package routes

import (
	"github.com/SerhiiCho/reciper/src/handlers"
	"github.com/SerhiiCho/reciper/src/utils"
	"github.com/gin-gonic/gin"
)

// LoadRoutes loads all available routes and serves the app
func LoadRoutes() {
	r := gin.Default()

	r.GET("/api/recipes", handlers.RecipesIndex())

	err := r.Run()
	utils.HandleError("Can't serve the app", err)
}
