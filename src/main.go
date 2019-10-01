package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// MyHandler struct
type MyHandler struct{}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := "app/dist/" + string(r.URL.Path[1:])

	log.Println(path)

	if path == "app/dist/" {
		path = "app/dist/index.html"
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404 My friend"))
		return
	}

	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml"
	} else {
		contentType = "text/html"
	}

	w.Header().Add("Content-Type", contentType)
	w.Write(data)
}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":3000", nil)
}
