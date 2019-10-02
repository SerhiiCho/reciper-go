package routes

import (
	"net/http"

	"github.com/SerhiiCho/reciper_go/src/handlers"
)

func handle(path string, controller func(http.ResponseWriter, http.Request)) {
	//
}

func loadRoutes() {
	handle("/recipes", handlers.RecipeIndexHandler)
}
