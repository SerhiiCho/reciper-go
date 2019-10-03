package controllers

import (
	"net/http"
)

// HomeController handler
func HomeController(w http.ResponseWriter, r *http.Request) {
	view(w, "home", map[string]interface{}{
		"title":    "Hello world",
		"array":    []string{"Hello", "nice", "it works", "other"},
		"isPretty": true,
	})
}
