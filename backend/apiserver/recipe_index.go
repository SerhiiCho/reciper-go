package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/SerhiiCho/reciper/backend/utils"
)

// recipeIndex handles GET request on showing the list of all recipes
func (serv *server) recipeIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipes := []map[string]string{
			{
				"title_ru":   "Cool title",
				"content_ru": "Cool content",
			},
			{
				"title_ru":   "Another title",
				"content_ru": "Some content",
			},
		}

		result, jsonErr := json.Marshal(recipes)
		utils.HandleError("error trying to marshal recipes map", jsonErr, "")

		_, err := w.Write([]byte(result))
		utils.HandleError("error trying to write response for recipes index", err, "")
	}
}
