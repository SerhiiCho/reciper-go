package cmd

import (
	"github.com/SerhiiCho/reciper/backend/api"
	"github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/app/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the api",
	Run: func(cmd *cobra.Command, args []string) {

		app := app.NewApp()
		api := api.NewAPI(app)

		defer app.Close()

		router := gin.Default().SecureJsonPrefix("/api")
		router.Use(getMiddlewares()...)

		api.Init(router)
	},
}

func getMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.AppMiddle(),
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
