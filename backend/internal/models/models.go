package models

import (
	"encoding/json"
	"time"
)

type Project struct {
	ID            int64     `json:"id"`
	Key           string    `json:"key"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	RetentionDays int64     `json:"retention_days"`
}

type Log struct {
	ID         int64            `json:"id"`
	ProjectKey string           `json:"project_key"`
	Timestamp  time.Time        `json:"timestamp"`
	Level      string           `json:"level"`
	Message    string           `json:"message"`
	AgentID    *string          `json:"agent_id"`
	Host       *string          `json:"host"`
	SourceFile *string          `json:"source_file"`
	Attrs      *json.RawMessage `json:"attrs"`
	IngestedAt time.Time        `json:"ingested_at"`
}

type Metric struct {
	ID         int64            `json:"id"`
	ProjectKey string           `json:"project_key"`
	Timestamp  time.Time        `json:"timestamp"`
	Name       string           `json:"name"`
	Value      float64          `json:"value"`
	MetricType string           `json:"metric_type"`
	AgentID    *string          `json:"agent_id"`
	Host       *string          `json:"host"`
	Tags       *json.RawMessage `json:"tags"`
	IngestedAt time.Time        `json:"ingested_at"`
}

type AlertRule struct {
	ID              int64            `json:"id"`
	ProjectKey      string           `json:"project_key"`
	Name            string           `json:"name"`
	Description     *string          `json:"description"`
	QueryType       string           `json:"query_type"`
	Filters         *json.RawMessage `json:"filters"`
	MetricName      *string          `json:"metric_name"`
	Operator        string           `json:"operator"`
	Threshold       float64          `json:"threshold"`
	WindowSeconds   int64            `json:"window_seconds"`
	Enabled         bool             `json:"enabled"`
	State           string           `json:"state"`
	LastTriggeredAt *time.Time       `json:"last_triggered_at"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type AlertEvent struct {
	ID         int64      `json:"id"`
	RuleID     int64      `json:"rule_id"`
	ProjectKey string     `json:"project_key"`
	FiredAt    time.Time  `json:"fired_at"`
	ResolvedAt *time.Time `json:"resolved_at"`
	Value      float64    `json:"value"`
	State      string     `json:"state"`
	Notified   bool       `json:"notified"`
}

type IngestLog struct {
	Timestamp  time.Time        `json:"timestamp"`
	Level      string           `json:"level"`
	Message    string           `json:"message"`
	AgentID    *string          `json:"agent_id,omitempty"`
	Host       *string          `json:"host,omitempty"`
	SourceFile *string          `json:"source_file,omitempty"`
	Attrs      *json.RawMessage `json:"attrs,omitempty"`
}

type IngestMetric struct {
	Timestamp  time.Time        `json:"timestamp"`
	Name       string           `json:"name"`
	Value      float64          `json:"value"`
	MetricType *string          `json:"metric_type,omitempty"`
	AgentID    *string          `json:"agent_id,omitempty"`
	Host       *string          `json:"host,omitempty"`
	Tags       *json.RawMessage `json:"tags,omitempty"`
}

type IngestBatch struct {
	Logs    []IngestLog    `json:"logs"`
	Metrics []IngestMetric `json:"metrics"`
}

type CreateProjectRequest struct {
	Key           string `json:"key"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	RetentionDays int64  `json:"retention_days"`
}
