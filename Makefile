start:
	go run -v ./...

build:
	docker build -t go-image:latest .

run:
	docker compose up

migrate-up:
	migrate -path internal/db/migrations/ -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up 1

migrate-down:
	migrate -path internal/db/migrations/ -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down 1