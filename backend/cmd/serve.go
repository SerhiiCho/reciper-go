package cmd

import (
	apiPackage "github.com/SerhiiCho/reciper/backend/api"
	appPackage "github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/app/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "serves the api",
		Run: func(cmd *cobra.Command, args []string) {

			app := appPackage.NewApp()
			api := apiPackage.NewAPI(app)

			router := gin.Default()
			router.Use(getMiddlewares()...)

			api.Init(router)
		},
	})
}

func getMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.AppMiddle(),
	}
}
