package rkn

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/config"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"time"
)

type RknBestClient struct {
	client *http.Client
	cfg    config.Rkn
}

func NewRknBestClient(cfg config.Rkn) *RknBestClient {
	return &RknBestClient{
		client: &http.Client{Timeout: time.Duration(cfg.DefaultTimeout) * time.Second},
		cfg:    cfg,
	}
}

func (c *RknBestClient) IsForbidden(srcUrl string) (bool, error) {
	srcHost, err := parseHost(srcUrl)
	if err != nil {
		return true, err
	}
	url := fmt.Sprintf(c.cfg.Url, srcHost)
	resp, err := http.Get(url)
	if err != nil {
		return true, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errMessage := fmt.Sprintf("Code: %d. Status: %s", resp.StatusCode, resp.Status)
		return true, errors.New(errMessage)
	}

	var siteInfo model.SiteInfo
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &siteInfo); err != nil {
		return true, errors.New("can not unmarshal JSON")
	}

	return siteInfo.Result != 1 || len(siteInfo.Data.List) == 0, nil
}

func parseHost(srcUrl string) (string, error) {
	u, err := netUrl.Parse(srcUrl)
	if err != nil {
		return "", err
	}

	return u.Hostname(), nil
}
