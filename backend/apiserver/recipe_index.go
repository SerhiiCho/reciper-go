package apiserver

import (
	"net/http"
)

// recipeIndex handles GET request on showing the list of all recipes
func (serv *server) recipeIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}
}
