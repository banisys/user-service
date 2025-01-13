new_migration:
	migrate create -ext sql -dir ./pkg/database/migrations -seq users

migrateup:
	migrate -path ./pkg/database/migrations -database "sqlite3://pkg/database/database.db" up