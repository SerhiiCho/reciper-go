package routes

import (
	"github.com/SerhiiCho/reciper/src/handlers"
	"net/http"
)

func init() {
	handle("recipes", handlers.RecipesIndexHandler)
}

func handle(path string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("/api/"+path, handler)
}
