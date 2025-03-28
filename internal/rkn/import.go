package rkn

import (
	"github.com/Koderbek/pocket_news_bot/internal/config"
	"github.com/Koderbek/pocket_news_bot/internal/repository"
	"time"
)

type Import struct {
	rknClient Client
	repo      *repository.Repository
	cfg       config.Import
}

func NewImport(rknClient Client, repo *repository.Repository, cfg config.Import) *Import {
	return &Import{rknClient: rknClient, repo: repo, cfg: cfg}
}

func (i *Import) Run() error {
	domains, err := i.rknClient.List()
	if err != nil {
		return err
	}

	var batch []string
	for _, domain := range domains {
		batch = append(batch, domain)
		if len(batch) == i.cfg.BatchSize {
			err = i.repo.DomainBlacklist.Save(batch)
			if err != nil {
				return err
			}

			batch = []string{}
			time.Sleep(i.cfg.DelayTime * time.Second)
		}
	}

	if len(batch) > 0 {
		err = i.repo.DomainBlacklist.Save(batch)
		if err != nil {
			return err
		}
	}

	return nil
}
