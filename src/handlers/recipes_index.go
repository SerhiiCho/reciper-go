package handlers

import (
	"encoding/json"
	"github.com/SerhiiCho/reciper/src/models"
	"github.com/SerhiiCho/reciper/src/utils"
	"net/http"
)

// RecipesIndex handles GET request on /api/recipes
func RecipesIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}

	var recipes []models.Recipe

	recipes = append(recipes, models.Recipe{
		ID:    1,
		Title: "Some title is here",
		Intro: "Intro will be at this position",
		Text:  "Text is also nice",
	})

	response, err := json.MarshalIndent(recipes, "", " ")
	utils.HandleError("Error while trying to marshal recipes", err)

	w.Header().Add("Content-Type", "application/json")

	_, writeErr := w.Write(response)
	utils.HandleError("Error writing to the response", writeErr)
}
