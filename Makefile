build:
	go build -o bin/student-crud-api cmd/main.go

run:
	go run cmd/main.go

migrate:
	export $(cat .env) && psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -f migrations/001_create_students_table.up.sql

test:
	go test ./...