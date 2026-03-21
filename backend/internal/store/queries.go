package store

import (
	"context"
	models "pulse/internal/models"
	"time"
)

func (s *Store) QueryProjects(ctx context.Context) ([]models.Project, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, key, name, description, created_at, retention_days, last_heartbeat FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]models.Project, 0)
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Key, &p.Name, &p.Description, &p.CreatedAt, &p.RetentionDays, &p.LastHeartbeat); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, rows.Err()
}

func (s *Store) InsertProject(ctx context.Context, key string, name string, description string, retentionDays int64) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx,
		"INSERT INTO projects (key, name, description, retention_days) VALUES (?,?,?,?) RETURNING id",
		key, name, description, retentionDays,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

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

func (s *Store) InsertLogs(ctx context.Context, projectKey string, logs []models.IngestLog) error {
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
			projectKey,
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

func (s *Store) InsertMetrics(ctx context.Context, projectKey string, metrics []models.IngestMetric) error {
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
			projectKey,
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
