package repository

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) GetAll() ([]model.Category, error) {
	var items []model.Category
	query := fmt.Sprintf("SELECT * FROM %s", categoryTable)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil
}
