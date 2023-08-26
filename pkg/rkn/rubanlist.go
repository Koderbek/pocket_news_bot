package rkn

import (
	"errors"
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/config"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const notExist = "Домена нет в реестре"

type RuBanListClient struct {
	client *http.Client
	cfg    config.Rkn
}

func NewRuBanListClientClient(cfg config.Rkn) *RuBanListClient {
	return &RuBanListClient{
		client: &http.Client{Timeout: cfg.DefaultTimeout * time.Second},
		cfg:    cfg,
	}
}

func (c *RuBanListClient) IsForbidden(srcUrl string) (bool, error) {
	srcHost, err := parseHost(srcUrl)
	if err != nil {
		return true, err
	}
	url := fmt.Sprintf(c.cfg.Url, srcHost)
	resp, err := http.Get(url)
	if err != nil {
		return true, err
	}

	if resp.StatusCode != http.StatusOK {
		errMessage := fmt.Sprintf("Code: %d. Status: %s", resp.StatusCode, resp.Status)
		return true, errors.New(errMessage)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return true, err
	}

	return !strings.Contains(string(body), notExist), nil
}
