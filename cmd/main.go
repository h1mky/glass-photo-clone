package main

import (
	"github.com/joho/godotenv"
	"glass-photo/internal/db"
	"glass-photo/internal/router"
	"log"
)

func main() {
	err := godotenv.Load("env")
	if err != nil {
		log.Fatalf("Error loading .env file,%s", err)
	}
	router.StartServer()
}
