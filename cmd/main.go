package main

import (
	"github.com/joho/godotenv"
	"glass-photo/internal/db"
	"glass-photo/internal/post"
	"glass-photo/internal/router"
	"glass-photo/internal/user"
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
	user.UserRouter()

	router.StartServer()

}
