package main

import (
	"github.com/BurntSushi/toml"
	"github.com/SerhiiCho/reciper/backend/apiserver"
	appPackage "github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("../.env")
	utils.HandleError("Error loading .env file", err, "")
}

func main() {
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
}
