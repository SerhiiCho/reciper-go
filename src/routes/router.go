package routes

import (
	"net/http"

	"handlers"
)

func handle(path string, controller func(http.ResponseWriter, http.Request)) {
	//
}

func loadRoutes() {
	handle("/recipes", handlers.RecipeIndexHandler)
}
