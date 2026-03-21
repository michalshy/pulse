package handler

import (
	"encoding/json"
	"net/http"
	"pulse/internal/models"
	"pulse/internal/store"
)

type Handler struct {
	store *store.Store
}

func New(store *store.Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) HandleBatch(w http.ResponseWriter, r *http.Request) {
	projectKey := r.Header.Get("Authorization")
	if projectKey == "" {
		http.Error(w, "Missing project key", http.StatusUnauthorized)
	}

	var batch models.IngestBatch
	if err := json.NewDecoder(r.Body).Decode(&batch); err != nil {
		http.Error(w, "Invalid body request", http.StatusBadRequest)
	}

	// batch exists
	if len(batch.Logs) > 0 {
		if err := h.store.InsertLogs(r.Context(), projectKey, batch.Logs); err != nil {
			http.Error(w, "Failed to insert logs", http.StatusInternalServerError)
		}
	}
	if len(batch.Metrics) > 0 {
		if err := h.store.InsertMetrics(r.Context(), projectKey, batch.Metrics); err != nil {
			http.Error(w, "Failed to insert metrics", http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
