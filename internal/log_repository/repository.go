package log_repository

import "database/sql"

type Logger interface {
	Error(source, message string) error
	Info(source, message string) error
}

type LogRepository struct {
	Logger
}

func NewLogRepository(db *sql.DB) (*LogRepository, error) {
	return &LogRepository{Logger: NewLogSqlite(db)}, nil
}
