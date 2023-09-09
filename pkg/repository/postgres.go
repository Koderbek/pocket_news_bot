package repository

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	categoryTable        = "category"
	chatCategoryTable    = "chat_category"
	sentNewsTable        = "sent_news"
	domainBlacklistTable = "domain_blacklist"
)

func NewPostgresDB(cfg config.Db) (*sqlx.DB, error) {
	connData := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", connData)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
