package main

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	logger2 "github.com/Koderbek/pocket_news_bot/internal/logger"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	"github.com/Koderbek/pocket_news_bot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	logger, err := logger2.Init("bot.log")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Init(false)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := repository.NewPostgresDB(cfg.Db)
	if err != nil {
		logger.Fatalf("Init DB error: %s", err.Error())
	}

	repos := repository.NewPostgresRepository(db)
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		logger.Fatal(err)
	}

	botApi.Debug = true
	bot := telegram.NewBot(botApi, repos, cfg.Messages)
	if err := bot.Start(); err != nil {
		logger.Fatal(err)
	}
}
