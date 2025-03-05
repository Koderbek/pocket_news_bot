package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChatCategoryPostgres_Add(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewChatCategoryPostgres(db)
	defer db.Close()

	type paramType struct {
		chatId     int64
		categoryId int8
		name       string
	}

	testCases := []struct {
		name       string
		param      paramType
		shouldFail bool
	}{
		{
			name:       "case-1: valid result",
			param:      paramType{1, 1, "TestChat"},
			shouldFail: false,
		},
		{
			name:       "case-2: valid result",
			param:      paramType{1, 2, "TestChat"},
			shouldFail: false,
		},
		{
			name:       "case-3: valid result - duplication",
			param:      paramType{1, 1, "TestChat"},
			shouldFail: false,
		},
		{
			name:       "case-4: return error - no category",
			param:      paramType{1, 111, "TestChat"},
			shouldFail: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.Add(tc.param.chatId, tc.param.categoryId, tc.param.name)
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestChatCategoryPostgres_Deactivate(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewChatCategoryPostgres(db)
	defer db.Close()

	type paramType struct {
		chatId     int64
		categoryId int8
	}

	testCases := []struct {
		name  string
		param paramType
	}{
		{
			name:  "case-1: successful deactivate",
			param: paramType{1, 2},
		},
		{
			name:  "case-2: no chat",
			param: paramType{111, 1},
		},
		{
			name:  "case-3: no category",
			param: paramType{1, 111},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.Deactivate(tc.param.chatId, tc.param.categoryId)
			assert.NoError(t, err)

			res := r.HasChatCategory(tc.param.chatId, tc.param.categoryId)
			assert.Equal(t, res, false)
		})
	}
}

func TestChatCategoryPostgres_GetByChatId(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewChatCategoryPostgres(db)
	defer db.Close()

	testCases := []struct {
		name   string
		param  int64
		result []model.ChatCategory
	}{
		{
			name:  "case-1: valid result",
			param: 1,
			result: []model.ChatCategory{
				{ChatId: 1, ChatName: "TestChat", CategoryId: 1, CategoryCode: "general", CategoryName: "Главное"},
			},
		},
		{
			name:   "case-2: empty result",
			param:  111,
			result: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.GetByChatId(tc.param)
			assert.NoError(t, err)
			assert.Equal(t, tc.result, got)
		})
	}
}

func TestChatCategoryPostgres_GetByCategoryId(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewChatCategoryPostgres(db)
	defer db.Close()

	result := []model.ChatCategory{
		{ChatId: 1, ChatName: "TestChat", CategoryId: 1, CategoryCode: "general", CategoryName: "Главное"},
	}

	testCases := []struct {
		name   string
		param  int8
		result []model.ChatCategory
	}{
		{
			name:   "case-1: valid result",
			param:  1,
			result: result,
		},
		{
			name:   "case-2: empty result - deactivated",
			param:  2,
			result: nil,
		},
		{
			name:   "case-3: empty result - no category",
			param:  111,
			result: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := r.GetByCategoryId(tc.param)
			assert.NoError(t, err)
			assert.Equal(t, tc.result, got)
		})
	}
}

func TestChatCategoryPostgres_HasChatCategory(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewChatCategoryPostgres(db)
	defer db.Close()

	type paramType struct {
		chatId     int64
		categoryId int8
	}

	testCases := []struct {
		name   string
		param  paramType
		result bool
	}{
		{
			name:   "case-1: find",
			param:  paramType{1, 1},
			result: true,
		},
		{
			name:   "case-2: not found - deactivated category",
			param:  paramType{1, 2},
			result: false,
		},
		{
			name:   "case-3: not found - no chat",
			param:  paramType{111, 1},
			result: false,
		},
		{
			name:   "case-4: not found - no category",
			param:  paramType{1, 111},
			result: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := r.HasChatCategory(tc.param.chatId, tc.param.categoryId)
			assert.Equal(t, tc.result, got)
		})
	}
}

func TestChatCategoryPostgres_DeactivateChat(t *testing.T) {
	godotenv.Load("../../.env") // Загружаем переменные окружения из .env
	cfg, _ := config.Init(true)
	db, _ := NewPostgresTestDB(cfg.Db)
	r := NewChatCategoryPostgres(db)
	defer db.Close()

	testCases := []struct {
		name  string
		param int64
	}{
		{
			name:  "case-1: successful deactivate",
			param: 1,
		},
		{
			name:  "case-2: no chat",
			param: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.DeactivateChat(tc.param)
			assert.NoError(t, err)

			res, err := r.GetByChatId(tc.param)
			assert.NoError(t, err)
			assert.Equal(t, res, []model.ChatCategory(nil))
		})
	}
}
