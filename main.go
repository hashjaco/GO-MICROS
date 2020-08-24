package main

import (
	"github.com/hashjaco/GO-MICROS/go-rest-api/app"
	"github.com/joho/godotenv"
	"log"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app.StartApplication()
}
