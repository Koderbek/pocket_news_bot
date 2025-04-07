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

	// Создаем срез интерфейсов для передачи в Exec
	args := make([]interface{}, len(domains))
	for i, domain := range domains {
		args[i] = domain
	}

	// Генерация плейсхолдеров ($1, $2, $3, ...)
	placeholders := make([]string, len(domains))
	for i := range domains {
		placeholders[i] = fmt.Sprintf("($%d)", i+1)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (domain) VALUES %s ON CONFLICT (domain) DO NOTHING;",
		DomainBlacklistTable,
		strings.Join(placeholders, ","),
	)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *DomainBlacklistPostgres) IsExists(searchDomain string) bool {
	var domain string
	query := fmt.Sprintf("SELECT domain FROM %s WHERE domain=$1", DomainBlacklistTable)
	err := r.db.Get(&domain, query, searchDomain)

	return err == nil && domain == searchDomain
}
