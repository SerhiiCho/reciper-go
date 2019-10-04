package routes

import (
	"github.com/SerhiiCho/reciper/src/handlers"
	"github.com/SerhiiCho/reciper/src/utils"
	"net/http"
	"os"
)

// LoadRoutes loads all available routes and serves the app
func LoadRoutes() {
	http.HandleFunc("/api/recipes", handlers.RecipesIndex)

	err := http.ListenAndServe(os.Getenv("APP_PORT"), nil)
	utils.HandleError("Listen and Serve error", err)
}
