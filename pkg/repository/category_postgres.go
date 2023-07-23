package repository

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
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
	query := fmt.Sprintf("SELECT * FROM %s", categoryTable)
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
