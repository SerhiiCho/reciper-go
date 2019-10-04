package routes

import (
	"github.com/SerhiiCho/reciper/src/utils"
	"net/http"
)

func setNotFoundPage(w http.ResponseWriter) {
	w.WriteHeader(404)
	_, err := w.Write([]byte("<h1>404 Page not found</h1>"))
	utils.HandleError("Error writing 404 page", err)
}
