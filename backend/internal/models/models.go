package models

import (
	"encoding/json"
	"time"
)

type Project struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	APIKey        string    `json:"api_key"`
	CreatedAt     time.Time `json:"created_at"`
	RetentionDays int64     `json:"retention_days"`
}

type Log struct {
	ID         string          `json:"id"`
	ProjectID  string          `json:"project_id"`
	Timestamp  time.Time       `json:"timestamp"`
	Level      string          `json:"level"`
	Message    string          `json:"message"`
	AgentID    string          `json:"agent_id"`
	Host       string          `json:"host"`
	SourceFile string          `json:"source_file"`
	Attrs      json.RawMessage `json:"attrs"`
	IngestedAt time.Time       `json:"ingested_at"`
}

type Metric struct {
	ID         string          `json:"id"`
	ProjectID  string          `json:"project_id"`
	Timestamp  time.Time       `json:"timestamp"`
	Name       string          `json:"name"`
	Value      float64         `json:"value"`
	MetricType string          `json:"metric_type"`
	AgentID    string          `json:"agent_id"`
	Host       string          `json:"host"`
	Tags       json.RawMessage `json:"tags"`
	IngestedAt time.Time       `json:"ingested_at"`
}

type AlertRule struct {
	ID              string          `json:"id"`
	ProjectID       string          `json:"project_id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	QueryType       string          `json:"query_type"`
	Filters         json.RawMessage `json:"filters"`
	MetricName      string          `json:"metric_name"`
	Operator        string          `json:"operator"`
	Threshold       float64         `json:"threshold"`
	WindowSeconds   int64           `json:"window_seconds"`
	Enabled         bool            `json:"enabled"`
	State           string          `json:"state"`
	LastTriggeredAt time.Time       `json:"last_triggered_at"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}
