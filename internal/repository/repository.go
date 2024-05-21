package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/jmoiron/sqlx"
)

type Category interface {
	GetAll() ([]model.Category, error)
	GetByCode(code string) (*model.Category, error)
	UpdateLastSent(code string) error
}

type ChatCategory interface {
	Create(chatId int64, categoryId int8, name string) error
	Delete(chatId int64, categoryId int8) error
	GetByChatId(chatId int64) ([]model.ChatCategory, error)
	GetByCategoryId(categoryId int8) ([]model.ChatCategory, error)
	HasChatCategory(chatId int64, categoryId int8) bool
}

type SentNews interface {
	Save(linksHash []string) error
	IsExists(linkHash string) bool
	Clean() error
}

type DomainBlacklist interface {
	Save(domains []string) error
	IsExists(domain string) bool
	IsEmpty() bool
}

type Repository struct {
	Category
	ChatCategory
	SentNews
	DomainBlacklist
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Category:        NewCategoryPostgres(db),
		ChatCategory:    NewChatCategoryPostgres(db),
		SentNews:        NewSentNewsPostgres(db),
		DomainBlacklist: NewDomainBlacklistPostgres(db),
	}
}
