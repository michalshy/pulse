package handler

import (
	"context"
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

func GetProject(ctx context.Context)
