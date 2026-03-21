package main

import (
	"log"
	"net/http"
	"os"
	"pulse/internal/handler"
	"pulse/internal/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	store, err := store.New("data/observability.go")
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	log.Println("Connected to database")

	h := handler.New(store)

	r := chi.NewRouter()
	log.Println("Created router")

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	}))
	log.Println("Setup cors")

	r.Post("/ingest", h.HandleBatch)
	log.Println("Registered handlers")

	port := os.Getenv("SERVER_PORT")
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
