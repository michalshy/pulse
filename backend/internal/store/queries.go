package store

import (
	"context"
	models "pulse/internal/models"
	"time"
)

func (s *Store) QueryProjects(ctx context.Context) ([]models.Project, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, key, name, description, api_key, created_at, retention_days FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]models.Project, 0)
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Key, &p.Name, &p.Description, &p.APIKey, &p.CreatedAt, &p.RetentionDays); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, rows.Err()
}

func (s *Store) InsertProject(ctx context.Context, key string, name string, description string, apiKey string, retentionDays int64) (int64, error) {
	var id int64
	err := s.db.QueryRowContext(ctx,
		"INSERT INTO projects (key, name, description, api_key, retention_days) VALUES (?,?,?,?,?) RETURNING id",
		key, name, description, apiKey, retentionDays,
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

func (s *Store) GetProjectByAPIKey(ctx context.Context, apiKey string) (*models.Project, error) {
	var project models.Project

	err := s.db.QueryRowContext(ctx,
		"SELECT id, key, name, description, api_key, created_at, retention_days FROM projects WHERE api_key = ?", apiKey,
	).Scan(&project.ID, &project.Key, &project.Name, &project.Description, &project.APIKey, &project.CreatedAt, &project.RetentionDays)

	if err != nil {
		return nil, err
	}

	return &project, nil
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
