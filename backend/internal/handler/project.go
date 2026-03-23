package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"pulse/internal/models"
)

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	result, err := h.store.QueryProjects(r.Context())
	if err != nil {
		slog.Error("There was an error returning the projects")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Info("Returning the projects")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	slog.Info("Creating new project")
	var req models.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invald requset body", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name required", http.StatusBadRequest)
		return
	}

	id, err := h.store.InsertProject(r.Context(), req.Key, req.Name, req.Description, req.RetentionDays)
	if err != nil {
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{
		"id": id,
	})
}
