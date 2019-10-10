package main

import (
	"github.com/BurntSushi/toml"
	"github.com/SerhiiCho/reciper/backend/apiserver"
	"github.com/SerhiiCho/reciper/backend/utils"
)

func main() {
	config := apiserver.NewConfig()
	_, configErr := toml.DecodeFile("config/server.toml", config)
	utils.FatalIfError("toml decode error in main@main", configErr)

	api := apiserver.NewAPIServer(config)
	apiErr := api.Start()

	utils.FatalIfError("error string api server in main@main", apiErr)
}
