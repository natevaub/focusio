golang-migrate/migrate 
1. Create a new migration file
migrate create -ext sql -dir db/migrations -seq create_sessions_table

2. Apply migrations (run UP)
migrate -path db/migrations -database "postgres://user:password@localhost:5432/db?sslmode=disable" up

3. Rollback last migration (run DOWN)
migrate -path db/migrations -database "postgres://user:password@localhost:5432/db?sslmode=disable" down 1

4. Check current version
migrate -path db/migrations -database "postgres://user:password@localhost:5432/db?sslmode=disable" version

5. Force to a specific version (⚠️ use with caution)
migrate -path db/migrations -database "postgres://user:password@localhost:5432/db?sslmode=disable" force 1

...
sqlc generate



sqlc