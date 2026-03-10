package models

import (
	"encoding/json"
	"time"
)

type BaseMessage struct {
	Type string `json:"type"`
}

type HandshakeMessage struct {
	Type     string          `json:"type"`
	GameID   string          `json:"game_id"`
	Metadata json.RawMessage `json:"metadata"`
}

type TriggerMessage struct {
	Type string `json:"type"`
}

type FlushPayload struct {
	Type      string          `json:"type"`
	SessionID int64           `json:"session_id"`
	Metrics   []MetricPayload `json:"metrics"`
	Events    []EventPayload  `json:"events"`
}

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
