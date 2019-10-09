package cmd

import (
	"github.com/BurntSushi/toml"
	"github.com/SerhiiCho/reciper/backend/apiserver"
	appPackage "github.com/SerhiiCho/reciper/backend/app"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "serves the api",
		Run: func(cmd *cobra.Command, args []string) {
			config := apiserver.NewConfig()
			_, configErr := toml.DecodeFile("config/server.toml", config)

			if configErr != nil {
				log.Fatal(configErr)
			}

			app := appPackage.NewApp()
			api := apiserver.NewAPIServer(app, config)
			apiErr := api.Start()

			defer app.Close()

			if apiErr != nil {
				log.Fatal(apiErr)
			}
		},
	})
}
