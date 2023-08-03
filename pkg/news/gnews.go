package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/config"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"github.com/Koderbek/pocket_news_bot/pkg/repository"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultTimeout = 5 * time.Second

type GNewsClient struct {
	client *http.Client
	repo   repository.Repository
	cfg    config.News
}

func NewGNewsClient(repo repository.Repository, cfg config.News) *GNewsClient {
	return &GNewsClient{
		client: &http.Client{Timeout: defaultTimeout},
		repo:   repo,
		cfg:    cfg,
	}
}

func (c *GNewsClient) GetNews(category string) ([]model.Article, error) {
	url := fmt.Sprintf(c.cfg.Url, category, c.cfg.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errMessage := fmt.Sprintf("Code: %d. Status: %s", resp.StatusCode, resp.Status)
		return nil, errors.New(errMessage)
	}

	var articles model.Articles
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, errors.New("can not unmarshal JSON")
	}

	return articles.Articles, nil
}
