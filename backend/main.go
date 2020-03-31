package main

import (
	_ "image/jpeg"

	"github.com/BurntSushi/toml"
	"github.com/SerhiiCho/reciper/backend/apiserver"
	"github.com/SerhiiCho/reciper/backend/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := apiserver.NewConfig()
	_, configErr := toml.DecodeFile("config/server.toml", config)

	utils.FatalIfError("toml decode error in main@main", configErr)
	utils.FatalIfError("error starting api server in main@main", apiserver.Start(config))
}
