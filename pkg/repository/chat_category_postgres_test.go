package repository

import (
	"errors"
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestChatCategoryPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewChatCategoryPostgres(db)

	type paramType struct {
		chatId     int64
		categoryId int8
		name       string
	}

	testCases := []struct {
		name         string
		param        paramType
		shouldFail   bool
		mockBehavior func()
	}{
		{
			name:       "case-1: valid result",
			shouldFail: false,
			param:      paramType{1, 1, "world"},
			mockBehavior: func() {
				query := fmt.Sprintf("INSERT INTO %s", chatCategoryTable)
				mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:       "case-1: return error",
			shouldFail: true,
			param:      paramType{1, 1, "world"},
			mockBehavior: func() {
				query := fmt.Sprintf("INSERT INTO %s", chatCategoryTable)
				mock.ExpectExec(query).WillReturnError(errors.New("return error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()
			err := r.Create(tc.param.chatId, tc.param.categoryId, tc.param.name)
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestChatCategoryPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewChatCategoryPostgres(db)

	type paramType struct {
		chatId     int64
		categoryId int8
	}

	testCases := []struct {
		name         string
		param        paramType
		shouldFail   bool
		mockBehavior func()
	}{
		{
			name:       "case-1: valid result",
			param:      paramType{1, 1},
			shouldFail: false,
			mockBehavior: func() {
				query := fmt.Sprintf("DELETE FROM %s", chatCategoryTable)
				mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:       "case-1: return error",
			param:      paramType{1, 1},
			shouldFail: true,
			mockBehavior: func() {
				query := fmt.Sprintf("DELETE FROM %s", chatCategoryTable)
				mock.ExpectExec(query).WillReturnError(errors.New("return error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()
			err := r.Delete(tc.param.chatId, tc.param.categoryId)
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestChatCategoryPostgres_GetByChatId(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewChatCategoryPostgres(db)
	result := []model.ChatCategory{
		{ChatId: 1, ChatName: "test", CategoryId: 1, CategoryCode: "world", CategoryName: "Мир"},
		{ChatId: 1, ChatName: "test", CategoryId: 2, CategoryCode: "sport", CategoryName: "Спорт"},
	}

	testCases := []struct {
		name         string
		param        int64
		result       []model.ChatCategory
		shouldFail   bool
		mockBehavior func(chatCategories []model.ChatCategory)
	}{
		{
			name:       "case-1: valid result",
			param:      1,
			result:     result,
			shouldFail: false,
			mockBehavior: func(chatCategories []model.ChatCategory) {
				query := fmt.Sprintf("SELECT (.+) FROM %s (.+);", chatCategoryTable)
				rows := sqlmock.NewRows([]string{"chat_id", "chat_name", "category_id", "category_code", "category_name"})
				for _, cat := range chatCategories {
					rows.AddRow(cat.ChatId, cat.ChatName, cat.CategoryId, cat.CategoryCode, cat.CategoryName)
				}

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:       "case-1: return error",
			param:      1,
			result:     result,
			shouldFail: true,
			mockBehavior: func(chatCategories []model.ChatCategory) {
				query := fmt.Sprintf("SELECT (.+) FROM %s (.+);", chatCategoryTable)
				mock.ExpectQuery(query).WillReturnError(errors.New("return error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(tc.result)
			got, err := r.GetByChatId(tc.param)
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.result, got)
			}
		})
	}
}

func TestChatCategoryPostgres_GetByCategoryId(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewChatCategoryPostgres(db)
	result := []model.ChatCategory{
		{ChatId: 1, ChatName: "test", CategoryId: 1, CategoryCode: "world", CategoryName: "Мир"},
		{ChatId: 2, ChatName: "test_2", CategoryId: 1, CategoryCode: "world", CategoryName: "Мир"},
	}

	testCases := []struct {
		name         string
		param        int8
		result       []model.ChatCategory
		shouldFail   bool
		mockBehavior func(chatCategories []model.ChatCategory)
	}{
		{
			name:       "case-1: valid result",
			param:      1,
			result:     result,
			shouldFail: false,
			mockBehavior: func(chatCategories []model.ChatCategory) {
				query := fmt.Sprintf("SELECT (.+) FROM %s (.+);", chatCategoryTable)
				rows := sqlmock.NewRows([]string{"chat_id", "chat_name", "category_id", "category_code", "category_name"})
				for _, cat := range chatCategories {
					rows.AddRow(cat.ChatId, cat.ChatName, cat.CategoryId, cat.CategoryCode, cat.CategoryName)
				}

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:       "case-1: return error",
			param:      1,
			result:     result,
			shouldFail: true,
			mockBehavior: func(chatCategories []model.ChatCategory) {
				query := fmt.Sprintf("SELECT (.+) FROM %s (.+);", chatCategoryTable)
				mock.ExpectQuery(query).WillReturnError(errors.New("return error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(tc.result)
			got, err := r.GetByCategoryId(tc.param)
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.result, got)
			}
		})
	}
}

func TestChatCategoryPostgres_HasChatCategory(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewChatCategoryPostgres(db)

	type paramType struct {
		chatId     int64
		categoryId int8
	}

	testCases := []struct {
		name         string
		param        paramType
		result       bool
		mockBehavior func()
	}{
		{
			name:   "case-1: find",
			param:  paramType{1, 1},
			result: true,
			mockBehavior: func() {
				query := fmt.Sprintf("SELECT category_id FROM %s", chatCategoryTable)
				rows := sqlmock.NewRows([]string{"category_id"})
				rows.AddRow(1)

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:   "case-1: not found",
			param:  paramType{1, 1},
			result: false,
			mockBehavior: func() {
				query := fmt.Sprintf("SELECT category_id FROM %s", chatCategoryTable)
				rows := sqlmock.NewRows([]string{"category_id"})
				rows.AddRow(0)

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:   "case-1: return error",
			param:  paramType{1, 1},
			result: false,
			mockBehavior: func() {
				query := fmt.Sprintf("SELECT category_id FROM %s", chatCategoryTable)
				mock.ExpectQuery(query).WillReturnError(errors.New("return error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()
			got := r.HasChatCategory(tc.param.chatId, tc.param.categoryId)
			assert.Equal(t, tc.result, got)
		})
	}
}
