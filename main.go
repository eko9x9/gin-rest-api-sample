package main

import (
	"log"

	"github.com/eko9x9/gin-rest-api-sample/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	server := &app.Routes{}
	server.StartServer()
}
