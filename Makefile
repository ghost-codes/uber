migrateup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/uber?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/uber?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/uber?sslmode=disable" -verbose down
	
migratedown1:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/uber?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

createMigration:
	migrate create -ext sql -dir db/migrations -seq $(name)

test:
	go test -v -cover ./...

server:
	go run server.go

.PHONY: migrateup migratedown migrateup1 migratedown1 sqlc test server
