package news

import "github.com/Koderbek/pocket_news_bot/pkg/model"

type Client interface {
	GetNews(category string) ([]model.Article, error)
}
