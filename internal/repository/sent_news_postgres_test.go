package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSentNewsPostgres_Save(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewSentNewsPostgres(db)
	defer db.Close()

	testCases := []struct {
		name       string
		param      []string
		shouldFail bool
	}{
		{
			name:       "case-1: valid result",
			shouldFail: false,
			param:      []string{"test1", "test2"},
		},
		{
			name:       "case-2: duplicate result",
			param:      []string{"test1", "test2"},
			shouldFail: false,
		},
		{
			name:       "case-3: empty param",
			param:      []string{},
			shouldFail: false,
		},
		{
			name:       "case-4: return error",
			param:      []string{strings.Repeat("test", 20)},
			shouldFail: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.Save(tc.param)
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSentNewsPostgres_IsExists(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewSentNewsPostgres(db)
	defer db.Close()

	testCases := []struct {
		name   string
		param  string
		result bool
	}{
		{
			name:   "case-1: find",
			param:  "test1",
			result: true,
		},
		{
			name:   "case-2: not found",
			param:  "test111",
			result: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := r.IsExists(tc.param)
			assert.Equal(t, tc.result, got)
		})
	}
}

func TestSentNewsPostgres_Clean(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewSentNewsPostgres(db)
	defer db.Close()

	testCases := []struct {
		name string
	}{
		{
			name: "case-1: valid result",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.Clean()
			assert.NoError(t, err)
			assert.Equal(t, r.IsExists("test1"), false)
			assert.Equal(t, r.IsExists("test2"), false)
		})
	}
}
