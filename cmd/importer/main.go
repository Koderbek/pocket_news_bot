package main

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/log_repository"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	"github.com/Koderbek/pocket_news_bot/internal/rkn"
	"log"
)

const logSource = "importer"

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
	rknClient := rkn.NewRoskomsvobodaClient(cfg.Rkn)
	rknImport := rkn.NewImport(rknClient, repos, cfg.Import)
	if err := rknImport.Run(); err != nil {
		logRep.Error(logSource, err.Error())
	} else {
		logRep.Info(logSource, "Import is completed")
	}
}
