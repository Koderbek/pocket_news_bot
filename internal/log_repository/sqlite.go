package log_repository

import (
	"database/sql"
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"log"
	_ "modernc.org/sqlite"
)

func NewSqliteDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.RootPath+"/logs/logs.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			source TEXT NOT NULL,
			level TEXT NOT NULL,
			message TEXT NOT NULL,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
