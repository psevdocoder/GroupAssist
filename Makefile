# Makefile
include .env

migrate:
	migrate -path ./database/migrations -database ${DB_URL} up

migrate-down:
	migrate -path ./database/migrations -database ${DB_URL} down

run:
	go run ./cmd/app/main.go

build:
	go build ./cmd/app/main.go

docker-build:
	docker compose build --no-cache