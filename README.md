# Install sqlc
```bash
brew install sqlc
```

# Migrate

## Install

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Migrate

```bash
go build . && ./mongosteen db create:migration {filename}

## or

migrate create -ext sql -dir config/migrations -seq create_users_table

```

## Run Migrate

```bash
go build . && ./mongosteen db migrate

## or

migrate -database "postgres://mongosteen:123456@localhost:5435/mongosteen_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" up

```

## Migrate Down
```bash

go build . && ./mongosteen db migrate:down

## or

migrate -database "postgres://mongosteen:123456@localhost:5435/mongosteen_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" down 1

```

## 测试

首先需要安装 MailHog 并运行：

```bash
go install github.com/mailhog/MailHog@v1.0.1 && MailHog
```

# 测试覆盖率

```bash
go test -coverprofile=coverage.out ./...

go tool cover -html=coverage.out -o coverage.html 
```

# 生成 struct

```bash
ctrl *SessionController mongosteen/internal/controller.Controller
```