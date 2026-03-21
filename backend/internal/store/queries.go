package store

import (
	models "pulse/internal/db"
	"time"
)

func (s *Store) QueryLogs(projectID string, from, to time.Time, level string) ([]models.Log, error) {
	return nil, nil
}

func (s *Store) QueryMetrics() {

}

func (s *Store) ListAlertRules() {

}

func (s *Store) ListAlertEvents() {

}

func (s *Store) CreateAlertRule() {

}

func (s *Store) GetProjectByAPIKey() {

}

func (s *Store) InsertLogs() {

}

func (s *Store) InsertMetrics() {

}
