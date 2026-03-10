package db

import (
	"context"
	"encoding/json"
)

func CreateSession(ctx context.Context, gameID string, metadata json.RawMessage) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO sessions (game_id, metadata) VALUES ($1, $2) RETURNING id",
		gameID, metadata,
	).Scan(&id)
	return id, err
}

func EndSession(ctx context.Context, sessionID int64) error {
	_, err := Pool.Exec(ctx,
		"UPDATE sessions SET ended_at = NOW() WHERE id = $1", sessionID)
	return err
}
