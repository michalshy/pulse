package db

import (
	"context"
	"encoding/json"
	"time"
)

func CreateEvent(ctx context.Context, captureID int64, recordedAt time.Time, name string, metadata json.RawMessage) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO events (capture_id, recorded_at, name, metadata) VALUES ($1, $2, $3, $4) RETURNING id",
		captureID, recordedAt, name, metadata,
	).Scan(&id)
	return id, err
}
