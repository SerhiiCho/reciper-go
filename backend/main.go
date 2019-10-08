package main

import (
	"github.com/SerhiiCho/reciper/backend/cmd"
	"github.com/SerhiiCho/reciper/backend/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	utils.HandleError("Error loading .env file", err, "")
}

func main() {
	cmd.Execute()
}
