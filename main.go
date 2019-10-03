package main

import (
	"github.com/SerhiiCho/reciper/src/bootstrap"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	bootstrap.IndexPage()
}
