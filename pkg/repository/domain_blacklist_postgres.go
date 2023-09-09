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

	query := fmt.Sprintf("INSERT INTO %s (domain) values %s", domainBlacklistTable, strings.Join(values, ","))
	_, err := r.db.Exec(query)

	return err
}

func (r *DomainBlacklistPostgres) IsExists(domain string) bool {
	var isExists int8
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE domain=$1", domainBlacklistTable)
	err := r.db.Get(&isExists, query, domain)

	return err == nil && isExists == 1
}

func (r *DomainBlacklistPostgres) Clean() error {
	_, err := r.db.Exec(fmt.Sprintf("TRUNCATE %s", domainBlacklistTable))

	return err
}
