package routes

import (
	"github.com/SerhiiCho/reciper/src/handlers"
	"github.com/SerhiiCho/reciper/src/utils"
	"net/http"
	"os"
)

// LoadRoutes loads all available routes and serves the app
func LoadRoutes() {
	router := routeHandler{}

	http.Handle("/", new(routeHandler))
	http.HandleFunc("/api/recipes", handlers.RecipesIndexHandler)

	err := http.ListenAndServe(os.Getenv("APP_PORT"), router)
	utils.HandleError("Listen and Serve error", err)
}
