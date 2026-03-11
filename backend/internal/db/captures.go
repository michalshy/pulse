package db

import "context"

func CreateCapture(ctx context.Context, sessionID int64) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO captures (session_id) VALUES ($1) RETURNING id",
		sessionID,
	).Scan(&id)
	return id, err
}
