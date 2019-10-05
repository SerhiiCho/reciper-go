package main

import (
	"github.com/SerhiiCho/reciper/src/routes"
	"github.com/SerhiiCho/reciper/src/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env", ".env.example")
	utils.HandleError("Error loading .env file", err)
}

func main() {
	routes.LoadRoutes()
}
