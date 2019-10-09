package apiserver

import (
	"io"
	"net/http"
)

// recipeIndex handles GET request on showing the list of all recipes
func (api *APIServer) recipeIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//ctx.JSON(http.StatusOK, api.App.Database.GetRecipes())
		io.WriteString(w, "Hello")
	}
}
