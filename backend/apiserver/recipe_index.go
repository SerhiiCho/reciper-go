package apiserver

import (
	"encoding/json"
	"github.com/SerhiiCho/reciper/backend/utils"
	"net/http"
)

// recipeIndex handles GET request on showing the list of all recipes
func (serv *server) recipeIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipes := []map[string]string{
			map[string]string{
				"title":   "Cool title",
				"content": "Cool content",
			},
		}
		result, _ := json.Marshal(recipes)
		_, err := w.Write([]byte(result))
		utils.HandleError("error trying to write response for recipes index", err, "")
	}
}
