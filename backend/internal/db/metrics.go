package db

import (
	"context"
	"encoding/json"
	"pulse/internal/models"
	"time"
)

func CreateMetric(ctx context.Context, captureID int64, recordedAt time.Time, name string, valueType models.MetricValueType, value json.RawMessage) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO metrics (capture_id, recorded_at, name, value_type, value) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		captureID, recordedAt, name, valueType, value,
	).Scan(&id)
	return id, err
}

func GetMetrics(ctx context.Context, captureID int64) ([]models.Metric, error) {
	rows, err := Pool.Query(ctx,
		"SELECT id, capture_id, recorded_at, name, value_type, value FROM metrics WHERE capture_id = $1",
		captureID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	metrics := make([]models.Metric, 0)
	for rows.Next() {
		var m models.Metric
		if err := rows.Scan(&m.ID, &m.CaptureID, &m.RecordedAt, &m.Name, &m.ValueType, &m.Value); err != nil {
			return nil, err
		}
		metrics = append(metrics, m)
	}

	return metrics, rows.Err()
}
