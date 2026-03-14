-- ============================================================
-- Pulse Database Schema
-- ============================================================

CREATE TABLE projects (
    id              BIGSERIAL   PRIMARY KEY,
    key             TEXT        NOT NULL UNIQUE,
    name            TEXT        NOT NULL,
    description     TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    metadata        JSONB       DEFAULT '{}'
);
CREATE INDEX idx_projects_created_at ON projects (created_at);

-- ── SESSIONS ─────────────────────────────────────────────────
-- Represents a single project run / playtesting session

CREATE TABLE sessions (
    id          BIGSERIAL       PRIMARY KEY,
    project_key TEXT            NOT NULL REFERENCES projects(key) ON DELETE CASCADE,
    started_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    ended_at    TIMESTAMPTZ,
    metadata    JSONB           DEFAULT '{}'        -- engine version, platform, build, map etc
);

CREATE INDEX idx_sessions_project_key ON sessions (project_key);

-- ── CAPTURES ─────────────────────────────────────────────────
-- Represents a single flush triggered by dashboard or SDK rule

CREATE TABLE captures (
    id          BIGSERIAL       PRIMARY KEY,
    session_id  BIGINT          NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    captured_at TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_captures_session_id ON captures (session_id);
CREATE INDEX idx_captures_captured_at ON captures (captured_at);


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
    value       JSONB               NOT NULL
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
    metadata    JSONB           DEFAULT '{}'  -- optional payload, e.g. { "ability": "dash", "damage": 42 }
);

CREATE INDEX idx_events_capture_id  ON events (capture_id);
CREATE INDEX idx_events_recorded_at ON events (recorded_at);
CREATE INDEX idx_events_name        ON events (name);
