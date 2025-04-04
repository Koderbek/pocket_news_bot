# Pocket News Bot

Данный бот предназначен для рассылки новостей в Telegram. Источником новостей является сервис новостей - [Gnews](https://gnews.io/)

## Instalation

Запуск бота происходит через `Docker`, поэтому необходимо его наличие на машине

### Prerequisites

Предварительно необходимо подготовить `.env` файл и положить его в корень проекта

Пример `.env` файла:

```
ENABLE_ALPINE_PRIVATE_NETWORKING=true
TZ=Europe/Moscow

TELEGRAM_API_KEY=your_bot_api_key
NEWS_API_KEY=your_gnews_api_key

DB_CONNECTION=postgresql://test:qwerty@postgres_db:5432/test_db?sslmode=disable

POSTGRES_USER=test
POSTGRES_PASSWORD=qwerty
POSTGRES_DB=test_db

ROOT_PATH=/usr/local/src
```
#### Пояснения по поводу `TELEGRAM_API_KEY` и `NEWS_API_KEY`

В `TELEGRAM_API_KEY` прописывается токен созданного вами бота через `@BotFather`

В `NEWS_API_KEY` прописывается `API Key`, который нужно сгенерировать в [Gnews](https://gnews.io/)

### Запуск в Docker
Для запуска бота в `Docker` необходимо запустить команду:
```
docker compose build && docker compose up
```

## Usage
После запуска бота появится приветственное сообщение, в котором предлагается выбрать интересующие новостные категории запуском команды `/editcategory`:

![img.png](readme_img.png)

После запуска команды `/editcategory` появится список кнопок с категориями. Необходимо выбрать интересующие новостные категории:

![img_1.png](readme_img_1.png)

После настройки категорий, через некоторое время, бот начнет присылать новости по выбранным ранее категориям:

![img_2.png](readme_img_2.png)

## Technologies

* [Go](https://go.dev/) - Go 1.22
* [Postgresql](https://www.postgresql.org/) - Database Postgres 14.8
* [SQLite](https://www.sqlite.org/) - Database for logs
* [Migrate](https://github.com/golang-migrate/migrate) - Database migrations written in Go