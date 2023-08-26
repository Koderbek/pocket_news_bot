package repository

import (
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Category interface {
	GetAll() ([]model.Category, error)
	GetByCode(code string) (*model.Category, error)
}

type ChatCategory interface {
	Create(chatId int64, categoryId int8, name string) error
	Delete(chatId int64, categoryId int8) error
	GetByChatId(chatId int64) ([]model.ChatCategory, error)
	GetByCategoryId(categoryId int8) ([]model.ChatCategory, error)
	HasChatCategory(chatId int64, categoryId int8) bool
}

type Repository struct {
	Category
	ChatCategory
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Category:     NewCategoryPostgres(db),
		ChatCategory: NewChatCategoryPostgres(db),
	}
}
