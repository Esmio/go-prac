# Migrate

## Install

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Migrate

```bash
migrate create -ext sql -dir config/migrations -seq create_users_table 
```

## Run Migrate

```bash
migrate -database "postgres://mongosteen:123456@localhost:5435/mongosteen_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" up
```