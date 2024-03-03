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
	if len(domains) == 0 || domains == nil {
		return nil
	}

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

func (r *DomainBlacklistPostgres) IsExists(searchDomain string) bool {
	var domain string
	query := fmt.Sprintf("SELECT domain FROM %s WHERE domain=$1", domainBlacklistTable)
	err := r.db.Get(&domain, query, searchDomain)

	return err == nil && domain == searchDomain
}

func (r *DomainBlacklistPostgres) IsEmpty() bool {
	var val int8
	query := fmt.Sprintf("SELECT 1 val FROM %s LIMIT 1;", domainBlacklistTable)
	err := r.db.Get(&val, query)

	return err == nil && val != 1
}
