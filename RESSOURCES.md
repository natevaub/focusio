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
BOILER PLATE SETUP
Simple and scalable file structure
├── Dockerfile
├── cmd
│   └── main.go
├── db
│   ├── generated
│   │   ├── db.go
│   │   ├── models.go
│   │   └── users.sql.go
│   ├── migrations
│   │   ├── 20250605212305_create_table_users.down.sql
│   │   └── 20250605212305_create_table_users.up.sql
│   ├── queries
│   │   └── users.sql
│   └── sqlc.yaml
├── go.mod
├── go.sum
└── internal
    ├── api
    └── db

create a sqlc.yaml
mkdir -p backend/migrations backend/queries
version: "2"
sql:
  - engine: "postgresql"
    queries: "./queries"
    schema: "./migrations"
    gen:
      go:
        package: "db"
        out: "./generated"
        sql_package: "pgx/v5"

