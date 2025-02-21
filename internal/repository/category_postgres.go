package repository

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db            *sqlx.DB
	allCategories []model.Category
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) GetAll() ([]model.Category, error) {
	if r.allCategories != nil {
		return r.allCategories, nil
	}

	var items []model.Category
	query := fmt.Sprintf("SELECT id, name, code FROM %s ORDER BY last_sent, id", categoryTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}

	r.allCategories = items
	return r.allCategories, nil
}

func (r *CategoryPostgres) GetByCode(code string) (*model.Category, error) {
	categories, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, cat := range categories {
		if cat.Code == code {
			return &cat, nil
		}
	}

	return nil, nil
}

func (r *CategoryPostgres) ForSending() (*model.Category, error) {
	categories, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	if len(categories) > 0 {
		return &categories[0], nil
	}

	return nil, nil
}

func (r *CategoryPostgres) UpdateLastSent(code string) error {
	query := fmt.Sprintf("UPDATE %s SET last_sent = NOW() WHERE code = $1;", categoryTable)
	_, err := r.db.Exec(query, code)
	r.allCategories = nil

	return err
}
