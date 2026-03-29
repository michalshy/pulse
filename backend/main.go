package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"pulse/internal/agent"
	"pulse/internal/config"
	"pulse/internal/handler"
	"pulse/internal/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	// Parsing the config
	configPath := flag.String("config", "../pulse/pulse.toml", "path to config file")
	flag.Parse()
	config, err := config.ParseConfig(*configPath)
	if err != nil {
		panic(err)
	}

	// Start and connect to db
	store, err := store.New("data/observability.db")
	if err != nil {
		slog.Error("Error connecting to database: ", err)
	}
	slog.Info("Connected to database")

	// Configure REST API
	h := handler.New(store)
	r := chi.NewRouter()
	slog.Info("Created router")
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	}))
	slog.Info("Setup cors")
	r.Post("/ingest", h.HandleBatch)
	r.Post("/project", h.CreateProject)
	r.Get("/projects", h.GetProjects)
	r.Post("/heartbeat/{project_key}", h.KeepAlive)
	slog.Info("Registered handlers")

	// Create and connect to PULSE Agent
	agent, err := agent.New(config.Agent)
	if err != nil {
		slog.Error("Failed to connect to agent", "error", err)
	} else {
		slog.Info("Connected to agent, sending test heartbeat...")
		if err := agent.Heartbeat(); err != nil {
			slog.Error("Heartbeat failed", "error", err)
		} else {
			slog.Info("Hearbeat success, agent is alive")
		}
	}

	// Listen on the port
	port := os.Getenv("SERVER_PORT")
	slog.Info("Server starting", "port", config.Backend.Port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		slog.Error(err.Error())
	}
}
