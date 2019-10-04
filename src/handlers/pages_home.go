package handlers

import (
	"github.com/SerhiiCho/reciper/src/utils"
	"github.com/SerhiiCho/reciper/src/utils/httputils"
	"io/ioutil"
	"net/http"
)

// PagesHome handles request/response on the home page
// it will load javascript application
func PagesHome(w http.ResponseWriter, r *http.Request) {
	path := "app/dist/" + r.URL.Path[1:]

	if path == "app/dist/" {
		path = "app/dist/index.html"
	}

	data, readErr := ioutil.ReadFile(path)

	if utils.HandleError("Error reading "+path, readErr) {
		httputils.SetNotFound(w)
		return
	}

	w.Header().Add("Content-Type", httputils.SetHeader(path))

	_, writeErr := w.Write(data)

	utils.HandleError("Error writing to response", writeErr)
}
