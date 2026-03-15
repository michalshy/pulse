package models

import (
	"encoding/json"
	"time"
)

// DATABASE MODELS

type Project struct {
	ID          int64           `json:"id"`
	Key         string          `json:"key"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
	Metadata    json.RawMessage `json:"metadata"`
	Active      bool            `json:"active"`
}

type Session struct {
	ID         int64           `json:"id"`
	ProjectKey string          `json:"project_key"`
	StartedAt  time.Time       `json:"started_at"`
	EndedAt    *time.Time      `json:"ended_at"`
	Metadata   json.RawMessage `json:"metadata"`
}

type Capture struct {
	ID         int64
	SessionID  int64
	CapturedAt time.Time
}

type MetricValueType string

const (
	MetricValueTypeFloat  MetricValueType = "float"
	MetricValueTypeInt    MetricValueType = "int"
	MetricValueTypeVec2   MetricValueType = "vec2"
	MetricValueTypeVec3   MetricValueType = "vec3"
	MetricValueTypeString MetricValueType = "string"
	MetricValueTypeJSON   MetricValueType = "json"
)

type Metric struct {
	ID         int64           `json:"id"`
	CaptureID  int64           `json:"capture_id"`
	RecordedAt time.Time       `json:"recorded_at"`
	Name       string          `json:"name"`
	ValueType  MetricValueType `json:"value_type"`
	Value      json.RawMessage `json:"value"`
}

type Event struct {
	ID         int64           `json:"id"`
	CaptureID  int64           `json:"capture_id"`
	RecordedAt time.Time       `json:"recorded_at"`
	Name       string          `json:"name"`
	Metadata   json.RawMessage `json:"metadata"`
}
