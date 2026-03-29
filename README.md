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
- **IPC Protocol**: gRPC + Protobuf

## Architecture

```
[Backend App] --gRPC--> [Agent] --HTTP batches--> [/ingest API]
      |                     ↑
      |                ring buffer
      +--spawns on startup--+
      +--heartbeat----------+

Backend Multi-writer: stdout + file + agent gRPC writer
```

### How It Works

1. **Backend auto-starts agent** on startup using configured binary path
2. **Backend connects** to agent via gRPC (local or remote, same code)
3. **Logs flow** through `slog` multi-writer to stdout, file, and agent simultaneously
4. **Agent buffers** messages in a ring buffer
5. **Agent flushes** to `/ingest` endpoint when:
   - Buffer reaches threshold (80% full)
   - Timer expires (every 10 seconds)
6. **Heartbeat** tracks backend health via unary RPC

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

Defined in `proto/pulse.proto`, compiled to Go (`gen/pulse/`) and Rust (`tonic-build` via `build.rs`):

```protobuf
syntax = "proto3";

package pulse;

option go_package = "./gen/pulse";

service Pulse {
  rpc IngestLogs    (stream LogEntry)      returns (Ack);
  rpc IngestMetrics (stream Metric)        returns (Ack);
  rpc Heartbeat     (HeartbeatRequest)     returns (HeartbeatResponse);
}

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

message HeartbeatRequest {
  int64 timestamp = 1;
}

message HeartbeatResponse {
  int64 timestamp = 1;
  bool ok = 2;
}

message Ack {
  bool success = 1;
}
```

## Integration

### Backend (Go)
- Parse `pulse.toml` configuration
- Spawn agent process if `auto_start = true`
- Connect via gRPC client to `host:port` (with retry loop on startup)
- Custom `slog.Handler` streams log entries via `IngestLogs`
- Multi-writer setup: `io.MultiWriter(stdout, file, agentWriter)`
- Heartbeat goroutine calls `Heartbeat` RPC periodically
- Reconnection logic logs "agent detached/reconnected"
- Code generation: `protoc` + `protoc-gen-go` + `protoc-gen-go-grpc`

### Agent (Rust)
- Parse same `pulse.toml` configuration
- gRPC server (`tonic`) listens on configured port
- Implement `Pulse` service trait: `ingest_logs`, `ingest_metrics`, `heartbeat`
- Store incoming entries in ring buffer (fixed capacity)
- Flush conditions:
  - Timer tick (every `flush_interval_secs`)
  - Buffer threshold reached (`flush_threshold * buffer_size`)
- Convert Protobuf → JSON and batch POST to `/ingest`
- Handle buffer overflow by dropping oldest messages
- Code generation: `tonic-build` in `build.rs` (runs automatically on `cargo build`)

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
- [x] Config parser for `pulse.toml`
- [x] Agent process spawner
- [ ] gRPC client connection (with startup retry)
- [ ] Custom `slog.Handler` for agent
- [ ] Multi-writer integration
- [ ] Heartbeat sender
- [ ] Reconnection logic

### Agent
- [x] Config parser
- [ ] gRPC server (`tonic`)
- [ ] Ring buffer
- [ ] Flush timer + threshold logic
- [ ] HTTP client for batch ingestion
- [ ] Graceful shutdown with buffer flush

### Proto
- [ ] `proto/pulse.proto` definition
- [ ] Go codegen setup (`protoc` + plugins)
- [ ] Rust codegen setup (`build.rs` + `tonic-build`)

## Roadmap

**Phase 1** (Current):
- gRPC communication (local or remote agent)
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
