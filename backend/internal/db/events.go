package db

import (
	"context"
	"encoding/json"
	"pulse/internal/models"
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

func GetEvents(ctx context.Context, captureID int64) ([]models.Event, error) {
	rows, err := Pool.Query(ctx,
		"SELECT id, capture_id, recorded_at, name, metadata FROM events WHERE capture_id = $1",
		captureID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]models.Event, 0)
	for rows.Next() {
		var e models.Event
		if err := rows.Scan(&e.ID, &e.CaptureID, &e.RecordedAt, &e.Name, &e.Metadata); err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, rows.Err()
}
