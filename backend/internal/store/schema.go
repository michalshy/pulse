package store

import _ "embed"

//go:embed schema.sql
var schemaSQL string

func (s *Store) initSchema() error {
	_, err := s.db.Exec(schemaSQL)
	return err
}
