package httputils

import (
	"github.com/SerhiiCho/reciper/src/utils"
	"net/http"
)

// SetNotFound sends status 404 and html with page not found title
func SetNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	_, err := w.Write([]byte("<h1>404 Page not found</h1>"))
	utils.HandleError("Error writing 404 page", err)
}
