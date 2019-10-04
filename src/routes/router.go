package routes

import (
	"github.com/SerhiiCho/reciper/src/utils"
	"io/ioutil"
	"net/http"
)

type routeHandler struct{}

func (_ routeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := "app/dist/" + request.URL.Path[1:]

	if path == "app/dist/" {
		path = "app/dist/index.html"
	}

	data, readErr := ioutil.ReadFile(path)

	if utils.HandleError("Error reading "+path, readErr) {
		setNotFoundPage(writer)
		return
	}

	writer.Header().Add("Content-Type", setHeader(path))

	_, writeErr := writer.Write(data)

	utils.HandleError("Error writing to response", writeErr)
}
