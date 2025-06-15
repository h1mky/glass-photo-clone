package main

import (
	"github.com/joho/godotenv"
	"glass-photo/internal/db"
	"glass-photo/internal/post"
	"glass-photo/internal/router"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file,%s", err)
	}

	db.Connect()
	defer db.DB.Close()

	post.PostRouter()

	router.StartServer()

}
