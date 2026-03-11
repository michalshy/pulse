-- ============================================================
-- Pulse Database Schema
-- ============================================================

-- ── SESSIONS ─────────────────────────────────────────────────
-- Represents a single project run / playtesting session

CREATE TABLE sessions (
    id          BIGSERIAL       PRIMARY KEY,
    project_id     TEXT            NOT NULL,
    started_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    ended_at    TIMESTAMPTZ,
    metadata    JSONB                                       -- engine version, platform, build, map etc
);

CREATE INDEX idx_sessions_project_id ON sessions (project_id);


-- ── CAPTURES ─────────────────────────────────────────────────
-- Represents a single flush triggered by dashboard or SDK rule

CREATE TABLE captures (
    id          BIGSERIAL       PRIMARY KEY,
    session_id  BIGINT          NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    captured_at TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_captures_session_id ON captures (session_id);
CREATE INDEX idx_captures_captured_at ON captures (captured_at);


-- ── METRIC VALUE TABLES ───────────────────────────────────────

CREATE TABLE metric_value_float (
    id          BIGSERIAL       PRIMARY KEY,
    value       FLOAT8          NOT NULL
);

CREATE TABLE metric_value_int (
    id          BIGSERIAL       PRIMARY KEY,
    value       INT8            NOT NULL
);

CREATE TABLE metric_value_vec2 (
    id          BIGSERIAL       PRIMARY KEY,
    x           FLOAT8          NOT NULL,
    y           FLOAT8          NOT NULL
);

CREATE TABLE metric_value_vec3 (
    id          BIGSERIAL       PRIMARY KEY,
    x           FLOAT8          NOT NULL,
    y           FLOAT8          NOT NULL,
    z           FLOAT8          NOT NULL
);

CREATE TABLE metric_value_string (
    id          BIGSERIAL       PRIMARY KEY,
    value       TEXT            NOT NULL
);

CREATE TABLE metric_value_json (
    id          BIGSERIAL       PRIMARY KEY,
    value       JSONB           NOT NULL
);


-- ── METRICS ──────────────────────────────────────────────────
-- One row per sampled data point inside a capture window
-- recorded_at = when the SDK sampled this inside the ring buffer
-- value_type  = which metric_value_* table to join
-- value_id    = FK into that table

CREATE TYPE metric_value_type AS ENUM ('float', 'int', 'vec2', 'vec3', 'string', 'json');

CREATE TABLE metrics (
    id          BIGSERIAL           PRIMARY KEY,
    capture_id  BIGINT              REFERENCES captures(id) ON DELETE CASCADE,
    recorded_at TIMESTAMPTZ         NOT NULL,
    name        TEXT                NOT NULL,
    value_type  metric_value_type   NOT NULL,
    value_id    BIGINT              NOT NULL
);

CREATE INDEX idx_metrics_capture_id ON metrics (capture_id);

-- ── EVENTS ───────────────────────────────────────────────────
-- Discrete named occurrences inside a capture window
-- e.g. 'boss_died', 'checkpoint_hit', 'ability_used'

CREATE TABLE events (
    id          BIGSERIAL       PRIMARY KEY,
    capture_id  BIGINT          NOT NULL REFERENCES captures(id) ON DELETE CASCADE,
    recorded_at TIMESTAMPTZ     NOT NULL,
    name        TEXT            NOT NULL,
    metadata    JSONB                                       -- optional payload, e.g. { "ability": "dash", "damage": 42 }
);

CREATE INDEX idx_events_capture_id  ON events (capture_id);
CREATE INDEX idx_events_recorded_at ON events (recorded_at);
CREATE INDEX idx_events_name        ON events (name);
