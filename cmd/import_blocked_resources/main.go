package main

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	logger2 "github.com/Koderbek/pocket_news_bot/internal/logger"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	"github.com/Koderbek/pocket_news_bot/internal/rkn"
	"log"
)

func main() {
	logger, err := logger2.Init("import_blocked_resources.log")
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
	rknClient := rkn.NewRoskomsvobodaClient(cfg.Rkn)
	rknImport := rkn.NewImport(rknClient, repos, cfg.Import)
	if err := rknImport.Run(); err != nil {
		logger.Fatal(err)
	} else {
		logger.Println("[SUCCESS] Import is completed")
	}
}
