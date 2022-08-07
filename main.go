package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/l-leniac-l/golang-signing-app/app/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err)
	}

	router := routes.SetupRouter()

	router.Run()
}
