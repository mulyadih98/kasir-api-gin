package main

import (
	"kasir-api-gin/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}
	server := app.CreateServer()
	server.Run(":" + os.Getenv("PORT"))
}
