package db

import "context"

func CreateCapture(ctx context.Context, sessionID int64, trigger string) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO captures (session_id, trigger) VALUES ($1, $2) RETURNING id",
		sessionID, trigger,
	).Scan(&id)
	return id, err
}
