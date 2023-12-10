.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

test_db_run:
	docker run --name=news_bot_db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d postgres:14.8-alpine

migrate_create:
	migrate create -ext sql -dir ./migrations $(name)

migration_up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' up

migration_down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' down