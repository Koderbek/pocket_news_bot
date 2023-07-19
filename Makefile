.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

test_db_run:
	docker run --name=news_bot_db -e POSTGRES_PASSWORD='qwerty' -p 5437:5432 -d postgres:14.8-alpine

migration_up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5437/postgres?sslmode=disable' up

migration_down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5437/postgres?sslmode=disable' down