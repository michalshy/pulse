package app

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"pulse/internal/agent"
	"pulse/internal/config"
	"pulse/internal/handler"
	"pulse/internal/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type App struct {
	config config.AppConfig
	store  *store.Store
	r      *chi.Mux

	agent *agent.Agent
}

func (app *App) parseConfig() {
	var err error
	configPath := flag.String("config", "../pulse/pulse.toml", "path to config file")
	flag.Parse()
	app.config, err = config.ParseConfig(*configPath)
	if err != nil {
		panic(err)
	}
}

func (app *App) dbConfig() {
	// Start and connect to db
	var err error
	app.store, err = store.New("data/observability.db")
	if err != nil {
		slog.Error("Error connecting to database: ", err)
	}
	slog.Info("Connected to database")
}

func (app *App) restConfig() {
	// Configure REST API
	h := handler.New(app.store)
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
}

func (app *App) pulseConfig() {
	// Create and connect to PULSE Agent
	var err error
	app.agent, err = agent.New(app.config.Agent)
	if err != nil {
		slog.Error("Failed to connect to agent", "error", err)
	}
}

func (app *App) logConfig() {
}

func (app *App) serveAgent(ctx context.Context) {
	go app.agent.StartHeartbeat(ctx, app.config.Agent.HeartbeatIntervalSecs)
}

func (app *App) Configure(ctx context.Context) {
	app.parseConfig()
	app.dbConfig()
	app.restConfig()

	// additionally, lets configure pulse agent
	app.pulseConfig()
	// now we can configure log, if there is a reason for us to log to agent
	app.logConfig()
}

func (app *App) Serve(ctx context.Context) {
	if app.agent != nil {
		app.serveAgent(ctx)
	}

	// Listen on the port
	port := app.config.Backend.Port
	host := app.config.Backend.Host
	slog.Info("Server starting", "port", app.config.Backend.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), app.r); err != nil {
		slog.Error(err.Error())
	}
}
