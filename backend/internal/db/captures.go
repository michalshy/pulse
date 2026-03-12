package db

import (
	"context"
	"pulse/internal/models"
)

func CreateCapture(ctx context.Context, sessionID int64) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO captures (session_id) VALUES ($1) RETURNING id",
		sessionID,
	).Scan(&id)
	return id, err
}

func GetCaptures(ctx context.Context, sessionID int64) ([]models.Capture, error) {
	rows, err := Pool.Query(ctx,
		"SELECT id, session_id, captured_at FROM captures WHERE session_id = $1",
		sessionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	captures := make([]models.Capture, 0)
	for rows.Next() {
		var c models.Capture
		if err := rows.Scan(&c.ID, &c.SessionID, &c.CapturedAt); err != nil {
			return nil, err
		}
		captures = append(captures, c)
	}

	return captures, rows.Err()
}
