package main

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/backup"
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/log_repository"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	"log"
)

const logSource = "backup"

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

	backupTables := []string{repository.ChatCategoryTable, repository.NewsSourcesTable}
	bu := backup.NewBackup(db, cfg, backupTables)
	if err := bu.Run(); err != nil {
		logRep.Error(logSource, err.Error())
		return
	}

	logRep.Info(logSource, "Backup is completed")
}
