package repository

import "database/sql"

type Logger interface {
	Log(source string, message string) error
}

type LogRepository struct {
	Logger
}

func NewLogRepository(db *sql.DB) *LogRepository {
	return &LogRepository{Logger: NewLogSqlite(db)}
}
