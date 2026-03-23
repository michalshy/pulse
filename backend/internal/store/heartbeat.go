package store

import (
	"context"
)

func (s *Store) KeepAlive(ctx context.Context, projectKey string) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE projects SET last_heartbeat = now() WHERE key = ?", projectKey,
	)
	return err
}
