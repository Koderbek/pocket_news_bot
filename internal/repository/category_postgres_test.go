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
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewCategoryPostgres(db)
	defer db.Close()

	testCases := []struct {
		name   string
		result int
	}{
		{
			name:   "case-1: valid result",
			result: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.GetAll()
			assert.NoError(t, err)
			assert.IsType(t, []model.Category{}, got)
			assert.Equal(t, tc.result, len(got))
		})
	}
}

func TestCategoryPostgres_GetByCode(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewCategoryPostgres(db)
	defer db.Close()

	testCases := []struct {
		name   string
		param  string
		result *model.Category
	}{
		{
			name:   "case-1: valid result",
			param:  "general",
			result: &model.Category{Id: 1, Name: "Главное", Code: "general"},
		},
		{
			name:   "case-2: empty result",
			param:  "test",
			result: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.GetByCode(tc.param)
			assert.NoError(t, err)
			assert.Equal(t, tc.result, got)
		})
	}
}

func TestCategoryPostgres_UpdateLastSent(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewCategoryPostgres(db)
	defer db.Close()

	testCases := []struct {
		name   string
		param  string
		result *model.Category
	}{
		{
			name:   "case-1: valid result",
			param:  "sports",
			result: &model.Category{Id: 7, Name: "Спорт", Code: "sports"},
		},
		{
			name:  "case-2: non-existing record",
			param: "1qa2ws3ed4rf5tg",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.UpdateLastSent(tc.param)
			assert.NoError(t, err)
			if tc.result != nil {
				cats, err := r.GetAll()
				assert.NoError(t, err)
				assert.Equal(t, tc.result, &cats[len(cats)-1])
			}
		})
	}
}

func TestCategoryPostgres_ForSending(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewCategoryPostgres(db)
	defer db.Close()

	testCases := []struct {
		name   string
		result *model.Category
	}{
		{
			name:   "case-1: valid result",
			result: &model.Category{Id: 1, Name: "Главное", Code: "general"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.ForSending()
			assert.NoError(t, err)
			assert.Equal(t, tc.result, got)
		})
	}
}
