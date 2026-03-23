package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) KeepAlive(w http.ResponseWriter, r *http.Request) {
	projectKey := chi.URLParam(r, "project_key")
	slog.Info("Heartbeat", "project", projectKey)
	err := h.store.KeepAlive(r.Context(), projectKey)
	if err != nil {
		slog.Error("Heartbeat failed", "error", err)
		http.Error(w, "Couldn't update project heartbeat", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
