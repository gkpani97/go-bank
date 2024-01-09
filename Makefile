postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4-alpine3.18

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-bank

dropdb:
	docker exec -it postgres15 dropdb go-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go-bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "$(CURDIR):/src" -w /src sqlc/sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server

