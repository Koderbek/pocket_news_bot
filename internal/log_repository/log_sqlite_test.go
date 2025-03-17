package log_repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogSqlite_Error(t *testing.T) {
	godotenv.Load("../../.env_test") // Загружаем переменные окружения
	cfg, _ := config.Init()

	logDb, err := NewSqliteDB(cfg)
	assert.NoError(t, err)
	logRep, err := NewLogRepository(logDb)
	assert.NoError(t, err)

	defer logDb.Close()

	type args struct {
		source  string
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case-1: valid result",
			args: args{source: "test", message: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logRep.Error(tt.args.source, tt.args.message)
			assert.NoError(t, err)
		})
	}
}

func TestLogSqlite_Info(t *testing.T) {
	godotenv.Load("../../.env_test") // Загружаем переменные окружения
	cfg, _ := config.Init()

	logDb, err := NewSqliteDB(cfg)
	assert.NoError(t, err)
	logRep, err := NewLogRepository(logDb)
	assert.NoError(t, err)

	defer logDb.Close()

	type args struct {
		source  string
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case-1: valid result",
			args: args{source: "test", message: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logRep.Info(tt.args.source, tt.args.message)
			assert.NoError(t, err)
		})
	}
}
