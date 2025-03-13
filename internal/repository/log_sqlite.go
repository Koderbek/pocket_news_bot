package repository

import (
	"database/sql"
)

type LogSqlite struct {
	db *sql.DB
}

func NewLogSqlite(db *sql.DB) *LogSqlite {
	return &LogSqlite{db: db}
}

func (l *LogSqlite) Log(source, message string) error {
	_, err := l.db.Exec(
		"INSERT INTO logs (source, message) VALUES ($1, $2)",
		source,
		message,
	)

	return err
}
