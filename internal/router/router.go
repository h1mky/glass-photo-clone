package router

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

var R *chi.Mux

func init() {
	R = chi.NewRouter()
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
