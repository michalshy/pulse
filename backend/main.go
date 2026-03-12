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

	restHandler := &rest.Handler{Manager: manager}

	r := chi.NewRouter()
	r.Get("/sessions/{game_id}", restHandler.GetSessions)
	r.Get("/session/{session_id}", restHandler.GetSession)
	r.Get("/captures/{session_id}", restHandler.GetCaptures)
	r.Get("/metrics/{capture_id}", restHandler.GetMetrics)
	r.Get("/events/{capture_id}", restHandler.GetEvents)
	r.Post("/trigger/{session_id}", restHandler.TriggerCapture)
	r.Get("/ws", wsHandler.HandleWS)

	port := os.Getenv("SERVER_PORT")
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
