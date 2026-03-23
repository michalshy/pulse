package main

import (
	"log/slog"
	"net/http"
	"os"
	"pulse/internal/handler"
	"pulse/internal/plog"
	"pulse/internal/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}
	logPath := os.Getenv("LOG_FILE")
	logger := plog.Logger{}
	err := logger.ConfigureLogger(logPath)
	if err != nil {
		slog.Error("Couldn't create logger")
	}
	defer logger.CloseLogger()

	store, err := store.New("data/observability.db")
	if err != nil {
		slog.Error("Error connecting to database: ", err)
	}
	slog.Info("Connected to database")

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

	port := os.Getenv("SERVER_PORT")
	slog.Info("Server starting", "port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		slog.Error(err.Error())
	}
}
