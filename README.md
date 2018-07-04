# Testing performance differences between postgres/cockroachdb direct to JSON VS Go

## Local setup

### Postgres

```bash
docker run -dit --name postgres -e POSTGRES_USER=bench -e POSTGRES_PASSWORD=bench -p 5432:5432 -d postgres:10.4-alpine
```

### Cockroachdb

```bash
docker run -dit --name cockroach -p 26257:26257 cockroachdb/cockroach:v2.0.3
```

## Migrations

```bash
go run ./cmd/migrations/main.go
```

## Create fake data, a lot (millions of rows)
