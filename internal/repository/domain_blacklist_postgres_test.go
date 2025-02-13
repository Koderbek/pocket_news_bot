package repository

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestDomainBlacklistPostgres_Save(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewDomainBlacklistPostgres(db)
	testCases := []struct {
		name         string
		param        []string
		shouldFail   bool
		mockBehavior func()
	}{
		{
			name:       "case-1: valid result",
			shouldFail: false,
			param:      []string{"test_1.ru", "test_2.ru"},
			mockBehavior: func() {
				query := fmt.Sprintf("INSERT INTO %s", domainBlacklistTable)
				mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:       "case-2: return error",
			shouldFail: true,
			param:      []string{"test_1.ru", "test_2.ru"},
			mockBehavior: func() {
				query := fmt.Sprintf("INSERT INTO %s", domainBlacklistTable)
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

func TestDomainBlacklistPostgres_IsExists(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	r := NewDomainBlacklistPostgres(db)
	testCases := []struct {
		name         string
		param        string
		result       bool
		mockBehavior func(searchDomain string)
	}{
		{
			name:   "case-1: find",
			param:  "test_1.ru",
			result: true,
			mockBehavior: func(searchDomain string) {
				query := fmt.Sprintf("SELECT domain FROM %s", domainBlacklistTable)
				rows := sqlmock.NewRows([]string{"domain"}).AddRow(searchDomain)

				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:   "case-2: not found",
			param:  "test_2.ru",
			result: false,
			mockBehavior: func(searchDomain string) {
				query := fmt.Sprintf("SELECT domain FROM %s", domainBlacklistTable)
				rows := sqlmock.NewRows([]string{"domain"}).AddRow("")
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
		{
			name:   "case-3: return error",
			param:  "test_3.ru",
			result: false,
			mockBehavior: func(searchDomain string) {
				query := fmt.Sprintf("SELECT domain FROM %s", domainBlacklistTable)
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
