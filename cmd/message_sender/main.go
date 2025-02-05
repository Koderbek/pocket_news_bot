package main

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/consumer"
	logger2 "github.com/Koderbek/pocket_news_bot/internal/logger"
	"github.com/Koderbek/pocket_news_bot/internal/news"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	logger, err := logger2.Init("message_sender.log")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Init()
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

	newsClient := news.NewGNewsClient(repos, cfg.News)
	cnsmr := consumer.NewConsumer(botApi, newsClient, repos, cfg.Consumer)
	if err := cnsmr.Start(); err != nil {
		logger.Fatal(err)
	} else {
		logger.Println("[SUCCESS] Sending is completed")
	}
}
