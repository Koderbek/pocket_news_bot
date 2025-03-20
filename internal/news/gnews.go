package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	"io"
	"net/http"
	"time"
)

type GNewsClient struct {
	client *http.Client
	repo   *repository.Repository
	cfg    config.News
}

func NewGNewsClient(repo *repository.Repository, cfg config.News) *GNewsClient {
	return &GNewsClient{
		client: &http.Client{Timeout: cfg.DefaultTimeout * time.Second},
		repo:   repo,
		cfg:    cfg,
	}
}

func (c *GNewsClient) GetNews(category string) ([]model.Article, error) {
	resp, err := http.Get(c.makeUrl(category))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errMessage := fmt.Sprintf("Code: %d. Status: %s", resp.StatusCode, resp.Status)
		return nil, errors.New(errMessage)
	}

	var articles model.Articles
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, errors.New("GetNews: can not unmarshal JSON")
	}

	return articles.Articles, nil
}

func (c *GNewsClient) makeUrl(category string) string {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return fmt.Sprintf(c.cfg.Url, category, date.Format(time.RFC3339), c.cfg.ApiKey)
}
