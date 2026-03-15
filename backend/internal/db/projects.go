package db

import (
	"context"
	"pulse/internal/models"
)

func GetProjects(ctx context.Context) ([]models.Project, error) {
	rows, err := Pool.Query(ctx, `
		SELECT id, key, name, description, created_at, metadata,
			EXISTS(SELECT 1 FROM sessions s WHERE s.project_key = p.key AND s.ended_at IS NULL) AS active
		FROM projects p
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]models.Project, 0)
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Key, &p.Name, &p.Description, &p.CreatedAt, &p.Metadata, &p.Active); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, rows.Err()
}

func CheckProject(ctx context.Context, projectKey string) error {
	var id int64
	err := Pool.QueryRow(ctx,
		"SELECT id FROM projects WHERE key = $1",
		projectKey,
	).Scan(&id)
	return err
}

func CreateProject(ctx context.Context, key string, name string, description string) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO projects (key, name, description) VALUES ($1, $2, $3) RETURNING id",
		key, name, description,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
