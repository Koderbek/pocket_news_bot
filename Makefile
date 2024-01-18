.PHONY:

include .env
export

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

compose_up:
	docker-compose up > cliInput.log

test_db_run:
	docker run --name=news_bot_db -e POSTGRES_PASSWORD='${POSTGRES_PASSWORD}' -p 5432:5432 -d postgres:14.8-alpine

migrate_create:
	migrate create -ext sql -dir ./migrations $(name)

migration_up:
	migrate -path ./migrations -database '${DB_CONNECTION}' up

migration_down:
	migrate -path ./migrations -database '${DB_CONNECTION}' down