FROM golang:1.22-bookworm

# DuckDB C library (needed by go-duckdb driver)
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Data directory for DuckDB
RUN mkdir -p /data

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/backend ./cmd/backend

EXPOSE 8080

CMD ["backend"]
