package rkn

import (
	"github.com/Koderbek/pocket_news_bot/pkg/repository"
	"time"
)

const batchSize = 1000
const startHour = 2
const dailySleep = 1

type Import struct {
	rknClient Client
	repo      *repository.Repository
}

func NewImport(rknClient Client, repo *repository.Repository) *Import {
	return &Import{rknClient: rknClient, repo: repo}
}

func (i *Import) Run() error {
	for {
		if time.Now().Hour() != startHour {
			time.Sleep(dailySleep * time.Hour)
			continue
		}

		if err := i.repo.DomainBlacklist.Clean(); err != nil {
			return err
		}

		domains, err := i.rknClient.List()
		if err != nil {
			return err
		}

		var batch []string
		for _, domain := range domains {
			batch = append(batch, domain)

			if len(batch) == batchSize {
				err = i.repo.DomainBlacklist.Save(batch)
				if err != nil {
					return err
				}
			}
		}

		if len(batch) > 0 {
			err = i.repo.DomainBlacklist.Save(batch)
			if err != nil {
				return err
			}
		}
	}
}
