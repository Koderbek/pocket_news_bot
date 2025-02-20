package main

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	logger2 "github.com/Koderbek/pocket_news_bot/internal/logger"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	"log"
)

func main() {
	logger, err := logger2.Init("clean_sent_news.log")
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
	if err := repos.SentNews.Clean(); err != nil {
		logger.Fatal(err)
	} else {
		logger.Println("[SUCCESS] Clean sent_news is completed")
	}
}
