package models

import (
	"encoding/json"
	"time"
)

type BaseMessage struct {
	Type string `json:"type"`
}

type HandshakeMessage struct {
	Type      string          `json:"type"`
	ProjectID string          `json:"project_id"`
	Metadata  json.RawMessage `json:"metadata"`
}

type TriggerMessage struct {
	Type string `json:"type"`
}

type FlushPayload struct {
	Type    string          `json:"type"`
	Metrics []MetricPayload `json:"metrics"`
	Events  []EventPayload  `json:"events"`
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

type MetricPayload struct {
	Name       string          `json:"name"`
	RecordedAt time.Time       `json:"recorded_at"`
	ValueType  MetricValueType `json:"value_type"`
	Value      json.RawMessage `json:"value"`
}

type EventPayload struct {
	Name       string          `json:"name"`
	RecordedAt time.Time       `json:"recorded_at"`
	Metadata   json.RawMessage `json:"metadata"`
}
