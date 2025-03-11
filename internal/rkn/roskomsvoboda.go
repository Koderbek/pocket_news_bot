package rkn

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"io"
	"net/http"
	"time"
)

type RoskomsvobodaClient struct {
	client *http.Client
	cfg    config.Rkn
}

func NewRoskomsvobodaClient(cfg config.Rkn) *RoskomsvobodaClient {
	return &RoskomsvobodaClient{
		client: &http.Client{Timeout: cfg.DefaultTimeout * time.Second},
		cfg:    cfg,
	}
}

func (c *RoskomsvobodaClient) List() ([]string, error) {
	resp, err := c.client.Get(c.cfg.Url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errMessage := fmt.Sprintf("Code: %d. Status: %s", resp.StatusCode, resp.Status)
		return nil, errors.New(errMessage)
	}

	var domains []string
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &domains); err != nil {
		return nil, errors.New("can not unmarshal JSON")
	}

	return domains, nil
}
