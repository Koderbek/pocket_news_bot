package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDomainBlacklistPostgres_Save(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewDomainBlacklistPostgres(db)
	defer db.Close()

	testCases := []struct {
		name       string
		param      []string
		shouldFail bool
	}{
		{
			name:       "case-1: valid result",
			param:      []string{"test_1.ru", "test_2.ru"},
			shouldFail: false,
		},
		{
			name: "case-1: duplicate result",
			param: []string{
				"test_1.ru",
				"test_2.ru",
			},
			shouldFail: false,
		},
		{
			name:       "case-3: empty param",
			param:      []string{},
			shouldFail: false,
		},
		{
			name: "case-1: error result",
			param: []string{
				"test_1.ru",
				strings.Repeat("test.ru", 100),
			},
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

func TestDomainBlacklistPostgres_IsExists(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewDomainBlacklistPostgres(db)
	defer db.Close()

	testCases := []struct {
		name   string
		param  string
		result bool
	}{
		{
			name:   "case-1: find",
			param:  "test_1.ru",
			result: true,
		},
		{
			name:   "case-2: not found",
			param:  "test_999.ru",
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
