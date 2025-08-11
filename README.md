# VCH Customer & Invoice App (Go + Fiber + Postgres)

## Quick Run (local)

```bash
export DATABASE_URL="postgres://vch:changeme@localhost:5432/vch?sslmode=disable"
export ADMIN_PASSWORD="changeme"
export TRACKING_BASE="https://vuachuyentien.com/tracking/"

# DB (Postgres local)
sudo -u postgres psql -c "CREATE USER vch WITH PASSWORD 'changeme';"
sudo -u postgres psql -c "CREATE DATABASE vch OWNER vch;"

# build & run
go mod tidy
go run ./cmd/server

# tests
go test ./...
```
