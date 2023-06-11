postgresinit:
	docker run --name pg-for-go-mongosteen -p 5435:5432 -e POSTGRES_USER=mongosteen -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=mongosteen_dev -e PGDATA=/var/lib/postgresql/data/pgdata -v pg-go-mongosteen-data:/var/lib/postgresql/data postgres:14

postgres:
	docker exec -it pg-for-go-mongosteen psql

createdb:
	docker exec -it pg-for-go-mongosteen createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it pg-for-go-mongosteen dropdb mongosteen_dev

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown