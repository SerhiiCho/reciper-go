package main

import (
	"net/http"

	"github.com/SerhiiCho/reciper_go/src/controllers"
)

func main() {
	http.HandleFunc("/", controllers.HomeController)
	http.ListenAndServe(":3000", nil)
}
