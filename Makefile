new_migration:
	migrate create -ext sql -dir ./pkg/database/migrations -seq users

migrate:
	migrate -path ./pkg/database/migrations -database "sqlite3://pkg/database/database.db" -verbose up

migratedown:
	migrate -path ./pkg/database/migrations -database "sqlite3://pkg/database/database.db" -verbose down

test run:
	go test ./test/... -v -count=1

dev:
	go run ./cmd/user-service