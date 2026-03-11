package main

import (
	"log"
	"net/http"
	"os"
	"pulse/internal/db"
	"pulse/internal/rest"
	"pulse/internal/session"
	"pulse/internal/ws"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	manager := session.NewManager()
	log.Println("Created manager")

	wsHandler := &ws.Handler{Manager: manager}
	log.Println("Created WS Handler")

	restHandler := &rest.Handler{}

	r := chi.NewRouter()
	r.Get("/sessions/{game_id}", restHandler.GetSessions)
	r.Get("/sessions/{game_id}/{session_id}", restHandler.GetSession)
	r.Get("/sessions/{game_id}/{session_id}/captures", restHandler.GetCaptures)
	r.Get("/sessions/{game_id}/{session_id}/metrics", restHandler.GetMetrics)
	r.Get("/sessions/{game_id}/{session_id}/events", restHandler.GetEvents)
	r.Post("/sessions/{game_id}/{session_id}/trigger", restHandler.PostCapture)

	http.HandleFunc("/ws", wsHandler.HandleWS)

	port := os.Getenv("SERVER_PORT")
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
