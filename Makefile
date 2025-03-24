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
	go build -o ./.bin/import_blocked_resources cmd/import_blocked_resources/main.go
	chmod +x ./.bin/import_blocked_resources

run_import: build_import
	echo "Running import_blocked_resources..."
	./.bin/import_blocked_resources

build_clean:
	mkdir -p ./.bin
	go build -o ./.bin/clean_sent_news cmd/clean_sent_news/main.go
	chmod +x ./.bin/clean_sent_news

run_clean: build_clean
	echo "Running clean_sent_news..."
	./.bin/clean_sent_news

build_sender:
	mkdir -p ./.bin
	go build -o ./.bin/message_sender cmd/message_sender/main.go
	chmod +x ./.bin/message_sender

run_sender: build_sender
	echo "Running message_sender..."
	./.bin/message_sender

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