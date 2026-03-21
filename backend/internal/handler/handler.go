package handler

import (
	"net/http"
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

}
