-- ============================================================
-- Observability Platform — DuckDB Schema
-- Run once on backend startup (all statements are idempotent)
-- ============================================================

-- ------------------------------------------------------------
-- Projects
-- Every other table references project_id.
-- api_key is what the agent sends in the Authorization header;
-- the backend maps it to a project_id on ingest.
-- ------------------------------------------------------------
CREATE TABLE IF NOT EXISTS projects (
    key         VARCHAR PRIMARY KEY,
    name        VARCHAR NOT NULL,
    api_key     VARCHAR NOT NULL UNIQUE,
    created_at  TIMESTAMP DEFAULT current_timestamp,
    retention_days INTEGER DEFAULT 30
);

-- ------------------------------------------------------------
-- Logs
-- One row per log line. attrs holds arbitrary key/value JSON
-- extracted by the agent (parsed from JSON, logfmt, etc.).
-- ------------------------------------------------------------
CREATE SEQUENCE IF NOT EXISTS logs_id_seq START 1;
CREATE TABLE IF NOT EXISTS logs (
    id INTEGER PRIMARY KEY DEFAULT nextval('logs_id_seq'),
    project_id  VARCHAR NOT NULL,
    timestamp   TIMESTAMP NOT NULL,
    level       VARCHAR,              -- debug, info, warn, error, fatal
    message     TEXT,
    agent_id    VARCHAR,
    host        VARCHAR,
    source_file VARCHAR,              -- which log file this came from
    attrs       JSON,                 -- arbitrary fields from the log line
    ingested_at TIMESTAMP DEFAULT current_timestamp
);

-- Primary query pattern: filter by project + time range + level
CREATE INDEX IF NOT EXISTS idx_logs_project_time
    ON logs (project_id, timestamp DESC);

CREATE INDEX IF NOT EXISTS idx_logs_level
    ON logs (project_id, level, timestamp DESC);

-- ------------------------------------------------------------
-- Metrics
-- Prometheus-style numeric data points.
-- tags holds label key/values (e.g. {"method":"GET","path":"/api"})
-- ------------------------------------------------------------
CREATE SEQUENCE IF NOT EXISTS metrics_id_seq START 1;
CREATE TABLE IF NOT EXISTS metrics (
    id          INTEGER PRIMARY KEY DEFAULT nextval('metrics_id_seq'),
    project_id  VARCHAR NOT NULL,
    timestamp   TIMESTAMP NOT NULL,
    name        VARCHAR NOT NULL,     -- e.g. http_request_duration_seconds
    value       DOUBLE NOT NULL,
    metric_type VARCHAR DEFAULT 'gauge',  -- gauge, counter, histogram
    agent_id    VARCHAR,
    host        VARCHAR,
    tags        JSON,                 -- label key/values
    ingested_at TIMESTAMP DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_metrics_project_name_time
    ON metrics (project_id, name, timestamp DESC);

-- ------------------------------------------------------------
-- Alert rules
-- Each rule belongs to a project and defines a SQL condition
-- that the alert engine evaluates every tick (30s).
-- ------------------------------------------------------------
CREATE SEQUENCE IF NOT EXISTS alert_rules_id_seq START 1;
CREATE TABLE IF NOT EXISTS alert_rules (
    id          INTEGER PRIMARY KEY DEFAULT next_val('alert_rules_id_seq'),
    project_id  VARCHAR NOT NULL,
    name        VARCHAR NOT NULL,
    description TEXT,
    -- what to check
    query_type  VARCHAR NOT NULL,     -- 'log_count' | 'metric_threshold'
    filters     JSON,                 -- e.g. {"level":"error","service":"api"}
    metric_name VARCHAR,              -- only for metric_threshold rules
    -- condition
    operator    VARCHAR NOT NULL,     -- gt, gte, lt, lte, eq
    threshold   DOUBLE NOT NULL,
    window_seconds INTEGER NOT NULL DEFAULT 300,  -- look-back window
    -- state
    enabled     BOOLEAN DEFAULT true,
    state       VARCHAR DEFAULT 'ok', -- ok, firing, resolved
    last_triggered_at TIMESTAMP,
    created_at  TIMESTAMP DEFAULT current_timestamp,
    updated_at  TIMESTAMP DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS idx_alert_rules_project
    ON alert_rules (project_id, enabled);

-- ------------------------------------------------------------
-- Alert events
-- One row every time a rule fires. Prevents re-notification
-- and gives a history of incidents.
-- ------------------------------------------------------------
CREATE SEQUENCE IF NOT EXISTS alert_events_id_seq START 1;
CREATE TABLE IF NOT EXISTS alert_events (
    id          INTEGER PRIMARY KEY DEFAULT next_val('alert_events_id_seq'),
    rule_id     INTEGER NOT NULL,
    project_id  VARCHAR NOT NULL,
    fired_at    TIMESTAMP NOT NULL,
    resolved_at TIMESTAMP,
    value       DOUBLE,               -- the value that triggered it
    state       VARCHAR DEFAULT 'firing', -- firing, resolved
    notified    BOOLEAN DEFAULT false
);

CREATE INDEX IF NOT EXISTS idx_alert_events_rule
    ON alert_events (rule_id, fired_at DESC);

-- ------------------------------------------------------------
-- Future: spans table for distributed tracing
-- Leaving this as a comment so the shape is documented
-- but not created until we actually build trace support.
-- ------------------------------------------------------------
-- CREATE TABLE IF NOT EXISTS spans (
--     id            VARCHAR PRIMARY KEY,
--     trace_id      VARCHAR NOT NULL,
--     parent_span_id VARCHAR,
--     project_id    VARCHAR NOT NULL,
--     timestamp     TIMESTAMP NOT NULL,
--     duration_us   BIGINT NOT NULL,     -- microseconds
--     service       VARCHAR,
--     operation     VARCHAR,
--     status        VARCHAR,             -- ok, error
--     attrs         JSON,
--     ingested_at   TIMESTAMP DEFAULT current_timestamp
-- );
