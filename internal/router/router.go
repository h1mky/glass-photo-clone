package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
)

var R *chi.Mux

func init() {
	R = chi.NewRouter()
	R.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
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
