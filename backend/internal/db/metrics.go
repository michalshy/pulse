package db

import (
	"context"
	"encoding/json"
	"pulse/internal/models"
	"time"
)

func CreateMetric(ctx context.Context, captureID int64, recordedAt time.Time, name string, valueType models.MetricValueType, valueID int64) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO metrics (capture_id, recorded_at, name, value_type, value_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		captureID, recordedAt, name, valueType, valueID,
	).Scan(&id)
	return id, err
}

func CreateMetricFloat(ctx context.Context, value float64) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO metric_value_float (value) VALUES ($1) RETURNING id",
		value,
	).Scan(&id)
	return id, err
}

func CreateMetricInt(ctx context.Context, value int64) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO metric_value_int (value) VALUES ($1) RETURNING id",
		value,
	).Scan(&id)
	return id, err
}

func CreateMetricVec2(ctx context.Context, x float64, y float64) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO metric_value_vec2 (x, y) VALUES ($1, $2) RETURNING id",
		x, y,
	).Scan(&id)
	return id, err
}

func CreateMetricVec3(ctx context.Context, x float64, y float64, z float64) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO metric_value_vec3 (x, y,z ) VALUES ($1, $2, $3) RETURNING id",
		x, y, z,
	).Scan(&id)
	return id, err
}
func CreateMetricString(ctx context.Context, value string) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO metric_value_string (value) VALUES ($1) RETURNING id",
		value,
	).Scan(&id)
	return id, err
}

func CreateMetricJSON(ctx context.Context, value json.RawMessage) (int64, error) {
	var id int64
	err := Pool.QueryRow(ctx,
		"INSERT INTO metric_value_json (value) VALUES ($1) RETURNING id",
		value,
	).Scan(&id)
	return id, err
}
