package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type SentNewsPostgres struct {
	db *sqlx.DB
}

func NewSentNewsPostgres(db *sqlx.DB) *SentNewsPostgres {
	return &SentNewsPostgres{db: db}
}

func (r *SentNewsPostgres) Save(linksHash []string) error {
	if len(linksHash) == 0 || linksHash == nil {
		return nil
	}

	var values []string
	for _, hash := range linksHash {
		values = append(values, fmt.Sprintf("('%s')", hash))
	}

	query := fmt.Sprintf("INSERT INTO %s (url_hash_sum) values %s ON CONFLICT (url_hash_sum) DO NOTHING", sentNewsTable, strings.Join(values, ","))
	_, err := r.db.Exec(query)

	return err
}

func (r *SentNewsPostgres) IsExists(linkHash string) bool {
	var result string
	query := fmt.Sprintf("SELECT url_hash_sum FROM %s WHERE url_hash_sum=$1", sentNewsTable)
	err := r.db.Get(&result, query, linkHash)

	return err == nil && result == linkHash
}

func (r *SentNewsPostgres) DeleteByDate(date time.Time) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE created_at < $1;", sentNewsTable)
	_, err := r.db.Exec(sql, date.Format(time.DateTime))

	return err
}
