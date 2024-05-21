package repository

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestSentNewsPostgres_Save(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewSentNewsPostgres(db)
	testCases := []struct {
		name         string
		param        []string
		shouldFail   bool
		mockBehavior func()
	}{
		{
			name:       "case-1: valid result",
			shouldFail: false,
			param:      []string{"test1111", "test2222"},
			mockBehavior: func() {
				query := fmt.Sprintf("INSERT INTO %s", sentNewsTable)
				mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:       "case-2: return error",
			shouldFail: true,
			param:      []string{"test1111", "test2222"},
			mockBehavior: func() {
				query := fmt.Sprintf("INSERT INTO %s", sentNewsTable)
				mock.ExpectExec(query).WillReturnError(errors.New("return error"))
			},
		},
		{
			name:         "case-3: empty param",
			shouldFail:   false,
			param:        []string{},
			mockBehavior: func() {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()
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
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewSentNewsPostgres(db)
	testCases := []struct {
		name         string
		param        string
		result       bool
		mockBehavior func(linkHash string)
	}{
		{
			name:   "case-1: find",
			param:  "test1111",
			result: true,
			mockBehavior: func(linkHash string) {
				query := fmt.Sprintf("SELECT url_hash_sum FROM %s", sentNewsTable)
				rows := sqlmock.NewRows([]string{"url_hash_sum"}).AddRow(linkHash)

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:   "case-2: not found",
			param:  "test2222",
			result: false,
			mockBehavior: func(linkHash string) {
				query := fmt.Sprintf("SELECT url_hash_sum FROM %s", sentNewsTable)
				rows := sqlmock.NewRows([]string{"url_hash_sum"}).AddRow("")
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:   "case-3: return error",
			param:  "test2222",
			result: false,
			mockBehavior: func(linkHash string) {
				query := fmt.Sprintf("SELECT url_hash_sum FROM %s", sentNewsTable)
				mock.ExpectQuery(query).WillReturnError(errors.New("return error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(tc.param)
			got := r.IsExists(tc.param)
			assert.Equal(t, tc.result, got)
		})
	}
}

func TestSentNewsPostgres_Clean(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewSentNewsPostgres(db)
	testCases := []struct {
		name         string
		shouldFail   bool
		mockBehavior func()
	}{
		{
			name:       "case-1: valid result",
			shouldFail: false,
			mockBehavior: func() {
				query := fmt.Sprintf("TRUNCATE %s", sentNewsTable)
				mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:       "case-2: return error",
			shouldFail: true,
			mockBehavior: func() {
				query := fmt.Sprintf("TRUNCATE %s", sentNewsTable)
				mock.ExpectExec(query).WillReturnError(errors.New("return error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()
			err := r.Clean()
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
