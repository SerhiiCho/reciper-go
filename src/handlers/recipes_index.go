package handlers

import (
	"github.com/SerhiiCho/reciper/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RecipesIndex handles GET request on /api/recipes
func RecipesIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		var recipes []models.Recipe

		c.Header("Access-Control-Allow-Origin", "http://localhost:8000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "604800")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Access-Control-Request-Method")
		c.Header("Access-Control-Allow-Methods", "GET, POST")

		recipes = append(recipes, models.Recipe{
			ID:          1,
			Title:       "Some long title is in here",
			Excerpt:     "Some title is here",
			Intro:       "Intro will be at this position",
			Text:        "Text is also nice",
			Ingredients: "afsdfasdf\nsadfsadfds\nsdafadsfadsf\n",
			Slug:        "recipe_slug",
			Time:        243,
			Image:       "https://www.bbcgoodfood.com/sites/default/files/recipe-collections/collection-image/2013/05/egg-fried-cauliflower-rice.jpg",
			Ready:       true,
			Approved:    true,
			Published:   true,
			Simple:      true,
			CreatedAt:   "10-10-2018 10:15:08",
			//User: models.User{
			//	ID:           1,
			//	Name:         "Mikel",
			//	Status:       "admin",
			//	Email:        "some@email.com",
			//	Username:     "master",
			//	XP:           2343,
			//	StreakDays:   3,
			//	Popularity:   244,
			//	Active:       true,
			//	StreakCheck:  "10-10-2018 10:15:08",
			//	NotifCheck:   "10-10-2018 10:15:08",
			//	ContactCheck: "10-10-2018 10:15:08",
			//	CreatedAt:    "10-10-2018 10:15:08",
			//	Photo:        "https://avatars3.githubusercontent.com/u/35465417?s=460&v=4",
			//},
		})

		c.JSON(http.StatusOK, recipes)
	}
}
