package main

import (
	"github.com/joho/godotenv"
	"go-library-manager/internal/database"
	"go-library-manager/internal/routes"
	"log"
	"os"
)

func main() {
	if os.Getenv("GO_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	database.Connect()
	routes.HandleRequests()
}
