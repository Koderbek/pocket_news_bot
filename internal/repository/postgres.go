package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	CategoryTable        = "category"
	ChatCategoryTable    = "chat_category"
	SentNewsTable        = "sent_news"
	NewsSourcesTable     = "news_sources"
	DomainBlacklistTable = "domain_blacklist"
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
