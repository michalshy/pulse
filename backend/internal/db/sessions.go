package db

import (
	"context"
	"encoding/json"
	"pulse/internal/models"
)

func CreateSession(ctx context.Context, projectID string, metadata json.RawMessage) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO sessions (project_id, metadata) VALUES ($1, $2) RETURNING id",
		projectID, metadata,
	).Scan(&id)
	return id, err
}

func EndSession(ctx context.Context, sessionID int64) error {
	_, err := Pool.Exec(ctx,
		"UPDATE sessions SET ended_at = NOW() WHERE id = $1", sessionID)
	return err
}

func GetSessions(ctx context.Context, projectID string) ([]models.Session, error) {
	rows, err := Pool.Query(ctx, "SELECT id, project_id, started_at, ended_at, metadata FROM sessions WHERE project_id = $1", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []models.Session
	for rows.Next() {
		var s models.Session
		if err := rows.Scan(&s.ID, &s.ProjectID, &s.StartedAt, &s.EndedAt, &s.Metadata); err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}

	return sessions, rows.Err()
}

func GetSession(ctx context.Context, sessionID int64) (*models.Session, error) {
	var s models.Session
	err := Pool.QueryRow(ctx,
		"SELECT id, project_id, started_at, ended_at, metadata FROM sessions WHERE id = $1",
		sessionID,
	).Scan(&s.ID, &s.ProjectID, &s.StartedAt, &s.EndedAt, &s.Metadata)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
