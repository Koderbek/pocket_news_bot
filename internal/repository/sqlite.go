package repository

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

func NewSqliteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./logs.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			source TEXT NOT NULL,
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
