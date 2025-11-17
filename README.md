# sentinel-core

**sentinel-core** — production-grade Go backend scaffold demonstrating Clean Architecture, REST API, Postgres, JWT auth skeleton, background worker pattern, Docker, and CI.

## Features
- Clean Architecture (internal/domain/usecase/infrastructure)
- REST API (chi)
- PostgreSQL (sqlx)
- Config via Viper
- Structured logging (zerolog)
- JWT auth skeleton
- Docker + docker-compose
- Example unit tests

## Quick start (development)
1. Copy `configs/config.yaml` and edit environment values.
2. Start postgres:
   ```bash
   docker-compose up -d
   ```
3. Run the server:
   ```bash
   go run ./cmd/api
   ```

## Project layout
See `/internal` for core app code and `/cmd/api` for app entrypoint.
