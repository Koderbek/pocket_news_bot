package repository

import (
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Category interface {
	GetAll() ([]model.Category, error)
}

type ChatCategory interface {
	Create(chatId, categoryId int, name string) error
	Delete(chatId, categoryId int) error
	GetByChatId(chatId int64) ([]model.ChatCategory, error)
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
