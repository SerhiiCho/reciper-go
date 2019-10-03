package bootstrap

import (
	"github.com/SerhiiCho/reciper/src/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// IndexPage loads first page
func IndexPage() {
	http.HandleFunc("/", loadStaticFiles)
	err := http.ListenAndServe(os.Getenv("APP_PORT"), nil)
	utils.HandleError("Listen and Serve error", err)
}

func loadStaticFiles(w http.ResponseWriter, r *http.Request) {
	path := "app/dist/" + r.URL.Path[1:]

	if path == "app/dist/" {
		path = "app/dist/index.html"
	}

	data, readErr := ioutil.ReadFile(path)

	if utils.HandleError("Error reading "+path, readErr) {
		setNotFoundPage(w)
		return
	}

	w.Header().Add("Content-Type", setHeader(path))
	_, writeErr := w.Write(data)

	utils.HandleError("Error writing to response", writeErr)
}

func setNotFoundPage(w http.ResponseWriter) {
	w.WriteHeader(404)
	_, err := w.Write([]byte("<h1>404 Page not found</h1>"))
	utils.HandleError("Error writing 404 page", err)
}

func setHeader(path string) string {
	switch {
	case strings.HasSuffix(path, ".css"):
		return "text/css"
	case strings.HasSuffix(path, ".png"):
		return "image/png"
	case strings.HasSuffix(path, ".js"):
		return "application/javascript"
	case strings.HasSuffix(path, ".svg"):
		return "image/svg+xml"
	default:
		return "text/html"
	}
}
