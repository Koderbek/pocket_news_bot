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

func TestCategoryPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := NewCategoryPostgres(db)
	result := []model.Category{
		{Id: 1, Name: "Главное", Code: "general"},
		{Id: 2, Name: "Мир", Code: "world"},
	}

	testCases := []struct {
		name         string
		result       []model.Category
		shouldFail   bool
		mockBehavior func(categories []model.Category)
	}{
		{
			name:       "case-1: return error",
			result:     result,
			shouldFail: true,
			mockBehavior: func(categories []model.Category) {
				query := fmt.Sprintf("SELECT id, name, code FROM %s ORDER BY last_sent, id", categoryTable)
				mock.ExpectQuery(query).WillReturnError(errors.New("return error"))
			},
		},
		{
			name:       "case-2: valid result",
			result:     result,
			shouldFail: false,
			mockBehavior: func(categories []model.Category) {
				query := fmt.Sprintf("SELECT id, name, code FROM %s ORDER BY last_sent, id", categoryTable)
				rows := sqlmock.NewRows([]string{"id", "name", "code"})
				for _, cat := range categories {
					rows.AddRow(cat.Id, cat.Name, cat.Code)
				}

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.result)
			got, err := r.GetAll()
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.result, got)
			}
		})
	}
}

func TestCategoryPostgres_GetByCode(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := NewCategoryPostgres(db)
	testCases := []struct {
		name         string
		param        string
		result       *model.Category
		shouldFail   bool
		mockBehavior func(code string, category *model.Category)
	}{
		{
			name:       "case-1: return error",
			param:      "general",
			result:     &model.Category{Id: 1, Name: "Главное", Code: "general"},
			shouldFail: true,
			mockBehavior: func(code string, category *model.Category) {
				query := fmt.Sprintf("SELECT id, name, code FROM %s ORDER BY last_sent, id", categoryTable)
				mock.ExpectQuery(query).WillReturnError(errors.New("return error"))
			},
		},
		{
			name:       "case-2: valid result",
			param:      "general",
			result:     &model.Category{Id: 1, Name: "Главное", Code: "general"},
			shouldFail: false,
			mockBehavior: func(code string, category *model.Category) {
				query := fmt.Sprintf("SELECT id, name, code FROM %s ORDER BY last_sent, id", categoryTable)
				rows := sqlmock.NewRows([]string{"id", "name", "code"}).AddRow(category.Id, category.Name, category.Code)

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:       "case-3: empty result",
			param:      "world",
			result:     nil,
			shouldFail: false,
			mockBehavior: func(code string, category *model.Category) {
				query := fmt.Sprintf("SELECT id, name, code FROM %s ORDER BY last_sent, id", categoryTable)
				rows := sqlmock.NewRows([]string{"id", "name", "code"}).AddRow(1, "Главное", "general")

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(tc.param, tc.result)
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

func TestCategoryPostgres_UpdateLastSent(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewCategoryPostgres(db)
	testCases := []struct {
		name         string
		param        string
		shouldFail   bool
		mockBehavior func()
	}{
		{
			name:       "case-1: OK",
			param:      "world",
			shouldFail: false,
			mockBehavior: func() {
				query := fmt.Sprintf("UPDATE %s SET", categoryTable)
				mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:       "case-1: return error",
			param:      "world",
			shouldFail: true,
			mockBehavior: func() {
				query := fmt.Sprintf("UPDATE %s SET", categoryTable)
				mock.ExpectQuery(query).WillReturnError(errors.New("return error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()
			err := r.UpdateLastSent(tc.param)
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
