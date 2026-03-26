package store

import (
	"context"
)

func (s *Store) KeepAlive(ctx context.Context, projectKey string) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE projects SET last_heartbeat = now() WHERE key = ?", projectKey,
	)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx,
		"INSERT INTO project_heartbeats (project_key) VALUES (?)", projectKey,
	)

	return err
}
