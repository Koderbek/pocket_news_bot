package repository

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

type NewsSourcesPostgres struct {
	db *sqlx.DB
}

func NewNewsSourcesPostgres(db *sqlx.DB) *NewsSourcesPostgres {
	return &NewsSourcesPostgres{db: db}
}

func (r *NewsSourcesPostgres) Save(newsSources []model.NewsSource) error {
	if len(newsSources) == 0 {
		return nil
	}

	//Очищаем данные от дублей
	newsSources = clearDuplicates(newsSources)

	// Создаем срез интерфейсов для передачи в Exec
	var args []interface{}
	for _, source := range newsSources {
		args = append(args, source.Domain, source.Category, source.Active)
	}

	// Генерация плейсхолдеров ($1, $2, $3, ...)
	placeholders := make([]string, len(newsSources))
	j := 1
	for i := 0; i < len(newsSources); i++ {
		placeholders[i] = fmt.Sprintf("($%d, $%d, $%d)", j, j+1, j+2)
		j = j + 3
	}

	query := fmt.Sprintf(`
			INSERT INTO %s (domain, category, active)
			VALUES %s
			ON CONFLICT (domain, category)
			DO UPDATE SET last_update = NOW(), active = EXCLUDED.active;
		`,
		NewsSourcesTable,
		strings.Join(placeholders, ","),
	)

	_, err := r.db.Exec(query, args...)
	return err
}

func clearDuplicates(sources []model.NewsSource) []model.NewsSource {
	var result []model.NewsSource
	keys := make(map[string]bool)
	for _, source := range sources {
		key := source.Domain + source.Category
		if _, ok := keys[key]; ok {
			continue
		}

		keys[key] = true
		result = append(result, source)
	}

	return result
}
