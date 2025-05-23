package repository

import (
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/jmoiron/sqlx"
	"time"
)

type Category interface {
	GetAll() ([]model.Category, error)
	GetByCode(code string) (*model.Category, error)
	ForSending() (*model.Category, error)
	UpdateLastSent(code string) error
}

type ChatCategory interface {
	Add(chatId int64, categoryId int8, name string) error
	Deactivate(chatId int64, categoryId int8) error
	DeactivateChat(chatId int64) error
	GetByChatId(chatId int64) ([]model.ChatCategory, error)
	GetByCategoryId(categoryId int8) ([]model.ChatCategory, error)
	HasChatCategory(chatId int64, categoryId int8) bool
}

type SentNews interface {
	Save(linksHash []string) error
	IsExists(linkHash string) bool
	DeleteByDate(date time.Time) error
}

type NewsSources interface {
	Save(newsSources []model.NewsSource) error
}

type DomainBlacklist interface {
	Save(domains []string) error
	IsExists(domain string) bool
}

type Repository struct {
	Category
	ChatCategory
	SentNews
	NewsSources
	DomainBlacklist
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Category:        NewCategoryPostgres(db),
		ChatCategory:    NewChatCategoryPostgres(db),
		SentNews:        NewSentNewsPostgres(db),
		NewsSources:     NewNewsSourcesPostgres(db),
		DomainBlacklist: NewDomainBlacklistPostgres(db),
	}
}
