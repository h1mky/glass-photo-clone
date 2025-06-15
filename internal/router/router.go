package router

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	R := chi.NewRouter()
	if err := http.ListenAndServe(":"+port, R); err != nil {
		log.Fatal(err)
	}

}
