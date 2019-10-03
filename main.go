package main

import (
	"github.com/SerhiiCho/reciper/src/routes"
	"github.com/SerhiiCho/reciper/src/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	utils.HandleError("Error loading .env file", err)
	routes.LoadRoutes()
}

func main() {
	//
}
