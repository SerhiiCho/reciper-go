package routes

import (
	"github.com/SerhiiCho/reciper/src/handlers"
	"net/http"
)

func handle(path string, controller func(http.ResponseWriter, http.Request)) {
	//
}

func loadRoutes() {
	handle("/recipes", handlers.RecipeIndexHandler)
}
