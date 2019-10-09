package main

import (
	"github.com/BurntSushi/toml"
	"github.com/SerhiiCho/reciper/backend/apiserver"
	appPackage "github.com/SerhiiCho/reciper/backend/app"
	"github.com/SerhiiCho/reciper/backend/db"
	"github.com/SerhiiCho/reciper/backend/utils"
	"io/ioutil"
)

func main() {
	apiConf, dbConf := getConfigs()

	database := db.NewDatabase(dbConf)
	app := appPackage.NewApp(database)
	api := apiserver.NewAPIServer(app, apiConf)
	apiErr := api.Start()

	defer app.Close()

	utils.FatalIfError(apiErr)
}

func getConfigs() (*apiserver.Config, *db.Config) {
	apiConf := apiserver.NewConfig()
	dbConf := db.NewConfig()

	config, readFileErr := ioutil.ReadFile("config/server.toml")

	utils.FatalIfError(readFileErr)

	_, apiDecodeErr := toml.Decode(string(config), apiConf)
	_, dbDecodeErr := toml.Decode(string(config), dbConf)

	utils.FatalIfError(apiDecodeErr)
	utils.FatalIfError(dbDecodeErr)

	return apiConf, dbConf
}
