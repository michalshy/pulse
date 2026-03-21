package store

import (
	"context"
	models "pulse/internal/models"
	"time"
)

func (s *Store) QueryLogs(projectID int64, from, to time.Time, level string) ([]models.Log, error) {
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

func (s *Store) GetProjectByAPIKey(ctx context.Context, apiKey string) (models.Project, error) {
	return models.Project{}, nil
}

func (s *Store) InsertLogs(ctx context.Context, projectID int64, logs []models.IngestLog) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for i := 0; i < len(logs); i++ {
		_, err = tx.ExecContext(ctx,
			`INSERT INTO logs 
			(project_id, timestamp, level, message, agent_id, host, source_file, attrs) 
			VALUES (?,?,?,?,?,?,?,?)`,
			projectID,
			logs[i].Timestamp,
			logs[i].Level,
			logs[i].Message,
			logs[i].AgentID,
			logs[i].Host,
			logs[i].SourceFile,
			logs[i].Attrs,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *Store) InsertMetrics(ctx context.Context, projectID int64, metrics []models.IngestMetric) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for i := 0; i < len(metrics); i++ {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO metrics 
			(project_id, timestamp, name, value, metric_type, agent_id, host, tags) 
			VALUES (?,?,?,?,?,?,?,?)`,
			projectID,
			metrics[i].Timestamp,
			metrics[i].Name,
			metrics[i].Value,
			metrics[i].MetricType,
			metrics[i].AgentID,
			metrics[i].Host,
			metrics[i].Tags,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
