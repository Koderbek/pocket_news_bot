package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryPostgres_GetAll(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init()
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewCategoryPostgres(db)
	defer db.Close()

	testCases := []struct {
		name       string
		result     int
		shouldFail bool
	}{
		{
			name:       "case-1: valid result",
			result:     9,
			shouldFail: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.GetAll()
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, []model.Category{}, got)
				assert.Equal(t, tc.result, len(got))
			}
		})
	}
}

func TestCategoryPostgres_GetByCode(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init()
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewCategoryPostgres(db)
	defer db.Close()

	testCases := []struct {
		name       string
		param      string
		result     *model.Category
		shouldFail bool
	}{
		{
			name:       "case-1: valid result",
			param:      "general",
			result:     &model.Category{Id: 1, Name: "Главное", Code: "general"},
			shouldFail: false,
		},
		{
			name:       "case-2: empty result",
			param:      "test",
			result:     nil,
			shouldFail: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.GetByCode(tc.param)
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.result, got)
			}
		})
	}
}
