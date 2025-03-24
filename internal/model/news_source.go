package model

import "time"

type NewsSource struct {
	Domain     string        `db:"domain"`
	Category   string        `db:"category"`
	Active     string        `db:"active"`
	LastUpdate time.Duration `db:"last_update"`
}
