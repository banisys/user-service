new_migration:
	migrate create -ext sql -dir ./pkg/database/migrations -seq users

migrate:
	migrate -path ./pkg/database/migrations -database "sqlite3://${DB_URL}" -verbose up

migratedown:
	migrate -path ./pkg/database/migrations -database "sqlite3://${DB_URL}" -verbose down