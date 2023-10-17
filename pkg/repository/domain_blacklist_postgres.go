package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DomainBlacklistPostgres struct {
	db *sqlx.DB
}

func NewDomainBlacklistPostgres(db *sqlx.DB) *DomainBlacklistPostgres {
	return &DomainBlacklistPostgres{db: db}
}

func (r *DomainBlacklistPostgres) Save(domains []string) error {
	var values []string
	for _, domain := range domains {
		values = append(values, fmt.Sprintf("('%s')", domain))
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (domain) VALUES %s ON CONFLICT (domain) DO NOTHING;",
		domainBlacklistTable,
		strings.Join(values, ","),
	)

	_, err := r.db.Exec(query)

	return err
}

func (r *DomainBlacklistPostgres) IsExists(domain string) bool {
	var isExists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE domain=$1)", domainBlacklistTable)
	err := r.db.Get(&isExists, query, domain)

	return err == nil && isExists
}

func (r *DomainBlacklistPostgres) IsEmpty() bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s);", domainBlacklistTable)
	err := r.db.Get(&exists, query)

	return err == nil && !exists
}
