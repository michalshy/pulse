# Pulse

A lightweight observability platform for aggregating logs, metrics, and traces across multiple projects.

## What It Does

Pulse provides centralized monitoring for your applications:
- **Logs**: Ingest structured logs with custom attributes
- **Metrics**: Collect time-series data (gauges, counters, histograms)
- **Alerts**: Configure rules based on log patterns or metric thresholds
- **Projects**: Isolate data per project with unique API keys
- **Heartbeat**: Track application health status

## Tech Stack

- **Backend**: Go + DuckDB (embedded analytics database)
- **Dashboard**: React + TypeScript + Vite + TailwindCSS
- **Agent**: Rust + Tokio (async runtime)
- **IPC Protocol**: TCP + Protobuf

## Architecture

```
[Backend App] --TCP+Protobuf--> [Agent] --HTTP batches--> [/ingest API]
      |                            ↑
      |                       ring buffer
      +--spawns on startup--------+
      +--heartbeat----------------+

Backend Multi-writer: stdout + file + agent socket
```

### How It Works

1. **Backend auto-starts agent** on startup using configured binary path
2. **Backend connects** to agent via TCP socket (platform-independent)
3. **Logs flow** through `slog` multi-writer to stdout, file, and agent simultaneously
4. **Agent buffers** messages in a ring buffer
5. **Agent flushes** to `/ingest` endpoint when:
   - Buffer reaches threshold (80% full)
   - Timer expires (every 10 seconds)
6. **Heartbeat** keeps connection alive and tracks backend health

**If agent is down**: Backend logs "agent detached" and continues with stdout + file only.

## Configuration

Single configuration file: **pulse.toml**

```toml
[project]
key = "pulse-backend"
name = "Pulse Backend"

[agent]
host = "127.0.0.1"
port = 9090
binary_path = "./agent/target/release/pulse-agent"
auto_start = true
buffer_size = 1000
flush_interval_secs = 10
flush_threshold = 0.8          # Flush at 80% capacity

[ingest]
endpoint = "http://localhost:8080/ingest"
timeout_secs = 5
retry_attempts = 3

[heartbeat]
interval_secs = 30
```

## Protobuf Schema

Messages sent over TCP use Protobuf for efficiency:

```protobuf
syntax = "proto3";

message LogEntry {
  int64 timestamp = 1;
  string level = 2;
  string message = 3;
  optional string agent_id = 4;
  optional string host = 5;
  optional string source_file = 6;
  map<string, string> attrs = 7;
}

message Metric {
  int64 timestamp = 1;
  string name = 2;
  double value = 3;
  optional string metric_type = 4;
  optional string agent_id = 5;
  optional string host = 6;
  map<string, string> tags = 7;
}

message Heartbeat {
  int64 timestamp = 1;
}

message Message {
  oneof payload {
    LogEntry log = 1;
    Metric metric = 2;
    Heartbeat heartbeat = 3;
  }
}
```

**Message Framing**: Length-prefixed (4 bytes u32 + protobuf payload)

## Integration

### Backend (Go)
- Parse `pulse.toml` configuration
- Spawn agent process if `auto_start = true`
- Connect to TCP socket at configured host:port
- Custom `slog.Handler` sends to agent via Protobuf
- Multi-writer setup: `io.MultiWriter(stdout, file, agentWriter)`
- Heartbeat goroutine sends periodic pings
- Reconnection logic logs "agent detached/reconnected"

### Agent (Rust)
- Parse same `pulse.toml` configuration
- TCP server listens on configured port
- Decode incoming Protobuf messages (using `prost` crate)
- Store in ring buffer (fixed capacity)
- Flush conditions:
  - Timer tick (every `flush_interval_secs`)
  - Buffer threshold reached (`flush_threshold * buffer_size`)
- Convert Protobuf → JSON and batch POST to `/ingest`
- Handle buffer overflow by dropping oldest messages

### Ring Buffer Behavior
- **Circular queue** with fixed capacity
- **Oldest messages dropped** when full
- Optionally emit "dropped N messages" metric

## Quick Start

**Backend**:
```bash
cd backend && go run main.go
```

**Dashboard**:
```bash
cd dashboard && npm install && npm run dev
```

**Agent** (auto-started by backend, or run manually):
```bash
cd agent && cargo build --release
cargo run
```

## Implementation Checklist

### Backend
- [ ] Config parser for `pulse.toml`
- [ ] Agent process spawner
- [ ] TCP client connection
- [ ] Protobuf encoder
- [ ] Custom `slog.Handler` for agent
- [ ] Multi-writer integration
- [ ] Heartbeat sender
- [ ] Reconnection logic

### Agent
- [ ] Config parser
- [ ] TCP server
- [ ] Protobuf decoder
- [ ] Ring buffer
- [ ] Flush timer + threshold logic
- [ ] HTTP client for batch ingestion
- [ ] Graceful shutdown with buffer flush

## Roadmap

**Phase 1** (Current):
- TCP+Protobuf communication
- Ring buffer with smart flushing
- Backend self-instrumentation
- Agent auto-start

**Phase 2** (Next):
- Log viewer with search and filtering
- Metrics visualization
- Alert engine with notifications
- Dashboard displays backend's own logs

**Phase 3** (Future):
- Multi-application support (agent collects from multiple apps)
- Distributed tracing
- Log file tailing
- System metrics collection

---

**Status**: Active Development | **Version**: 0.1.0 Alpha
