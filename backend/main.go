package main

import (
	"github.com/BurntSushi/toml"
	"github.com/SerhiiCho/reciper/backend/apiserver"
	"github.com/SerhiiCho/reciper/backend/store"
	"github.com/SerhiiCho/reciper/backend/utils"
	"io/ioutil"
)

func main() {
	apiConf, dbConf := getConfigs()

	db := store.NewStore(dbConf)
	api := apiserver.NewAPIServer(apiConf)
	apiErr := api.Start()

	defer db.Close()

	utils.FatalIfError(apiErr)
}

func getConfigs() (*apiserver.Config, *store.Config) {
	apiConf := apiserver.NewConfig()
	dbConf := store.NewConfig()

	config, readFileErr := ioutil.ReadFile("config/server.toml")

	utils.FatalIfError(readFileErr)

	_, apiDecodeErr := toml.Decode(string(config), apiConf)
	_, dbDecodeErr := toml.Decode(string(config), dbConf)

	utils.FatalIfError(apiDecodeErr)
	utils.FatalIfError(dbDecodeErr)

	return apiConf, dbConf
}
