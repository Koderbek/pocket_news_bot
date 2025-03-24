package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	categoryTable        = "category"
	chatCategoryTable    = "chat_category"
	sentNewsTable        = "sent_news"
	newsSourcesTable     = "news_sources"
	domainBlacklistTable = "domain_blacklist"
)

func NewPostgresDB(cfg config.Db) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.ConnectionUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
