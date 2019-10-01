package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SerhiiCho/reciper_go/src/bootstrap"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	http.HandleFunc("/", bootstrap.LoadStaticFiles)
	http.ListenAndServe(os.Getenv("APP_PORT"), nil)
}
