package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewsSourcesPostgres_Save(t *testing.T) {
	godotenv.Load("../../.env_test") // Загружаем переменные окружения
	cfg, _ := config.Init()
	db, _ := NewPostgresDB(cfg.Db)
	r := NewNewsSourcesPostgres(db)
	defer db.Close()

	testCases := []struct {
		name       string
		param      []model.NewsSource
		shouldFail bool
	}{
		{
			name:       "case-1: valid result",
			param:      []model.NewsSource{{Domain: "test.com", Category: "general", Active: "Y"}, {Domain: "test2.com", Category: "general", Active: "Y"}},
			shouldFail: false,
		},
		{
			name:       "case-1: duplicate result",
			param:      []model.NewsSource{{Domain: "test2.com", Category: "general", Active: "N"}},
			shouldFail: false,
		},
		{
			name:       "case-3: empty param",
			param:      []model.NewsSource{},
			shouldFail: false,
		},
		{
			name:       "case-1: error result",
			param:      []model.NewsSource{{Domain: "test3.com", Category: "general", Active: "Z"}},
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
