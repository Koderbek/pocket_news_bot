package main

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	"github.com/Koderbek/pocket_news_bot/internal/rkn"
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
	rknClient := rkn.NewRoskomsvobodaClient(cfg.Rkn)
	rknImport := rkn.NewImport(rknClient, repos, cfg.Import)
	if err := rknImport.Run(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("[SUCCESS] Import is completed")
	}
}
