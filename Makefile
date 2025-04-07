.PHONY:

#include .env_test
include .env
export

build_bot:
	mkdir -p ./.bin
	go build -o ./.bin/bot cmd/bot/main.go
	chmod +x ./.bin/bot

run_bot: build_bot
	echo "Running bot..."
	./.bin/bot

build_import:
	mkdir -p ./.bin
	go build -o ./.bin/importer cmd/importer/main.go
	chmod +x ./.bin/importer

run_import: build_import
	echo "Running importer..."
	./.bin/importer

build_clean:
	mkdir -p ./.bin
	go build -o ./.bin/cleaner cmd/cleaner/main.go
	chmod +x ./.bin/cleaner

run_clean: build_clean
	echo "Running cleaner..."
	./.bin/cleaner

build_backup:
	mkdir -p ./.bin
	go build -o ./.bin/backup cmd/backup/main.go
	chmod +x ./.bin/backup

run_backup: build_backup
	echo "Running backup..."
	./.bin/backup

build_sender:
	mkdir -p ./.bin
	go build -o ./.bin/sender cmd/sender/main.go
	chmod +x ./.bin/sender

run_sender: build_sender
	echo "Running sender..."
	./.bin/sender

compose_up:
	docker compose build
	docker compose up > cliInput.log

run_test:
	go test ./...

test_db_run:
	docker run --name=${POSTGRES_DB} -e POSTGRES_DB=${POSTGRES_DB_TEST} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -p 5433:5432 -d postgres:14.8-alpine

migrate_create:
	migrate create -ext sql -dir ./migrations $(name)

migration_up:
	migrate -path ./migrations -database '${DB_CONNECTION}' up

migration_down:
	migrate -path ./migrations -database '${DB_CONNECTION}' down