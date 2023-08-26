package main

import (
	"github.com/Koderbek/pocket_news_bot/pkg/config"
	"github.com/Koderbek/pocket_news_bot/pkg/consumer"
	"github.com/Koderbek/pocket_news_bot/pkg/news"
	"github.com/Koderbek/pocket_news_bot/pkg/repository"
	"github.com/Koderbek/pocket_news_bot/pkg/rkn"
	"github.com/Koderbek/pocket_news_bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(cfg.Db)
	if err != nil {
		log.Fatalf("Init DB error: %s", err.Error())
	}

	repos := repository.NewPostgresRepository(db)
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	botApi.Debug = true
	bot := telegram.NewBot(botApi, repos, cfg.Messages)

	newsClient := news.NewGNewsClient(repos, cfg.News)
	rknClient := rkn.NewRuBanListClientClient(cfg.Rkn)
	cnsmr := consumer.NewConsumer(botApi, newsClient, rknClient, repos, cfg.Consumer)
	go func() {
		if err := cnsmr.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
