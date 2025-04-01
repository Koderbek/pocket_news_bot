package main

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/log_repository"
	"github.com/Koderbek/pocket_news_bot/internal/message_sender"
	"github.com/Koderbek/pocket_news_bot/internal/news"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const logSource = "sender"

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	logDb, err := log_repository.NewSqliteDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer logDb.Close()

	logRep, err := log_repository.NewLogRepository(logDb)
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(cfg.Db)
	if err != nil {
		logRep.Error(logSource, fmt.Sprintf("Init DB error: %s", err.Error()))
	}
	defer db.Close()

	repos := repository.NewPostgresRepository(db)
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		logRep.Error(logSource, err.Error())
	}

	newsClient := news.NewGNewsClient(repos, cfg.News)
	sender := message_sender.NewSender(botApi, newsClient, repos, cfg.MessageSender)
	if err := sender.Start(); err != nil {
		logRep.Error(logSource, err.Error())
	} else {
		logRep.Info(logSource, "Sending is completed")
	}
}
