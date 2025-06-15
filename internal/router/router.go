package router

import (
	"github.com/go-chi/chi/v5"
	"glass-photo/internal/post"
	"log"
	"net/http"
	"os"
)

var R *chi.Mux

func init() {
	R = chi.NewRouter()
	initRoutes()
}

func initRoutes() {
	R.Get("/posts", post.GetAllPostHandler)       // route for getting all posts
	R.Get("/posts/{id}", post.GetPostByIDHandler) // route for getting a post by ID
}

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s...\n", port)

	if err := http.ListenAndServe(":"+port, R); err != nil {
		log.Fatal(err)
	}
}
