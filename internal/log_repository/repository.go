package log_repository

import (
	"database/sql"
	"time"
)

type Logger interface {
	Error(source, message string) error
	Info(source, message string) error
	DeleteByDate(date time.Time) error
}

type LogRepository struct {
	Logger
}

func NewLogRepository(db *sql.DB) (*LogRepository, error) {
	return &LogRepository{Logger: NewLogSqlite(db)}, nil
}
