package news

import "github.com/Koderbek/pocket_news_bot/internal/model"

type Client interface {
	GetNews(category string) ([]model.Article, error)
}
