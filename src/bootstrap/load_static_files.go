package bootstrap

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// IndexPage loads first page
func IndexPage() {
	loadEnv()
	http.HandleFunc("/", loadStaticFiles)
	http.ListenAndServe(os.Getenv("APP_PORT"), nil)
}

func loadStaticFiles(w http.ResponseWriter, r *http.Request) {
	path := "app/dist/" + string(r.URL.Path[1:])

	if path == "app/dist/" {
		path = "app/dist/index.html"
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		setNotFoundPage(w)
		return
	}

	w.Header().Add("Content-Type", setHeader(path))
	w.Write(data)
}

func setNotFoundPage(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Write([]byte("<h1>404 Page not found</h1>"))
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
