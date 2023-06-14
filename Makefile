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

mysqlinit:
	docker run -d --name mysql-for-go-mongosteen -p 3307:3306 -e MYSQL_DATABASE=mongosteen_dev -e MYSQL_USER=mongosteen -e MYSQL_PASSWORD=123456 -e MYSQL_ROOT_PASSWORD=123456 -v mysql-go-mongosteen-data:/var/lib/mysql mysql:8 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

mysql:
	docker exec -it mysql-for-go-mongosteen mysql -u mongosteen -p mongosteen_dev

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown mysqlinit