package models

import (
	"encoding/json"
	"time"
)

// DATABASE MODELS

type Session struct {
	ID int64
	GameID string
	StartedAt time.Time
	EndedAt *time.Time
	Metadata json.RawMessage
}

type Capture struct {
	ID int64
	SessionID int64
	CapturedAt time.Time
	Trigger string
}

type MetricValueFloat struct {
	ID int64
	Value float64 
}

type MetricValueInt struct {
	ID int64
	Value int64
}

type MetricValueVec2 struct {
	ID int64
	X float64
	Y float64
}

type MetricValueVec3 struct {
	ID int64
	X float64
	Y float64
	Z float64
}

type MetricValueJSON struct {
	ID int64
	Value json.RawMessage
}

type MetricValueText struct {
	ID int64
	Value string
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
	ID int64
	CaptureID *int64
	RecordedAt time.Time
	Name string
	ParentID *int64
	ValueType MetricValueType
	ValueID int64
}

type Event struct {
	ID int64
	CaptureID int64
	RecordedAt time.Time
	Name string
	Metadata json.RawMessage
}