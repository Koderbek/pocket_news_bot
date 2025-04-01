package log_repository

import (
	"database/sql"
	"fmt"
	"time"
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
		"INSERT INTO logs (source, level, message, timestamp) VALUES ($1, $2, $3, $4)",
		source,
		level,
		message,
		time.Now().Format(time.DateTime),
	)

	return err
}

func (l *LogSqlite) Error(source, message string) error {
	return l.log(source, errorLevel, message)
}

func (l *LogSqlite) Info(source, message string) error {
	return l.log(source, infoLevel, message)
}

func (l *LogSqlite) DeleteByDate(date time.Time) error {
	_, err := l.db.Exec(
		fmt.Sprintf("DELETE FROM logs WHERE timestamp < $1;"),
		date.Format(time.DateTime),
	)

	return err
}
