package main

import (
	"github.com/BurntSushi/toml"
	"github.com/SerhiiCho/reciper/backend/apiserver"
	appPackage "github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/db"
	"io/ioutil"
	"log"
)

func main() {
	apiConf, dbConf := getConfigs()

	database := db.NewDatabase(dbConf)
	app := appPackage.NewApp(database)
	api := apiserver.NewAPIServer(app, apiConf)
	apiErr := api.Start()

	defer app.Close()

	if apiErr != nil {
		log.Fatal(apiErr)
	}
}

func getConfigs() (*apiserver.Config, *db.Config) {
	apiConf := apiserver.NewConfig()
	dbConf := db.NewConfig()

	config, readFileErr := ioutil.ReadFile("config/server.toml")

	if readFileErr != nil {
		log.Fatal(readFileErr)
	}

	_, apiDecodeErr := toml.Decode(string(config), apiConf)
	_, dbDecodeErr := toml.Decode(string(config), dbConf)

	if apiDecodeErr != nil {
		log.Fatal(apiDecodeErr)
	}

	if dbDecodeErr != nil {
		log.Fatal(dbDecodeErr)
	}

	return apiConf, dbConf
}
