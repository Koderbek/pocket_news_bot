package log_repository

import (
	"database/sql"
)

const errorLevel = "ERROR"
const infoLevel = "INFO"

type LogSqlite struct {
	db *sql.DB
}

func NewLogSqlite(db *sql.DB) *LogSqlite {
	return &LogSqlite{db: db}
}

func (l *LogSqlite) log(source, level, message string) error {
	_, err := l.db.Exec(
		"INSERT INTO logs (source, level, message) VALUES ($1, $2, $3)",
		source,
		level,
		message,
	)

	return err
}

func (l *LogSqlite) Error(source, message string) error {
	return l.log(source, errorLevel, message)
}

func (l *LogSqlite) Info(source, message string) error {
	return l.log(source, infoLevel, message)
}
