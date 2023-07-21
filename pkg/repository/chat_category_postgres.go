package repository

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/pkg/model"
	"github.com/jmoiron/sqlx"
)

type ChatCategoryPostgres struct {
	db *sqlx.DB
}

func NewChatCategoryPostgres(db *sqlx.DB) *ChatCategoryPostgres {
	return &ChatCategoryPostgres{db: db}
}

func (r *ChatCategoryPostgres) Create(chatId, categoryId int, name string) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (chat_id, category_id, name) values ($1, $2, $3)",
		chatCategoryTable,
	)

	_, err := r.db.Exec(query, chatId, categoryId, name)

	return err
}

func (r *ChatCategoryPostgres) Delete(chatId, categoryId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE chat_id = $1 AND category_id = $2",
		chatCategoryTable,
	)

	_, err := r.db.Exec(query, chatId, categoryId)
	return err
}

func (r *ChatCategoryPostgres) GetByChatId(chatId int64) ([]model.ChatCategory, error) {
	var items []model.ChatCategory
	query := fmt.Sprintf(`
		SELECT cc.chat_id,
			   cc.name as chat_name,
			   cc.category_id,
			   c.code  as category_code,
			   c.name  as category_name
		FROM %s cc
				 INNER JOIN %s c on c.id = cc.category_id
		WHERE chat_id = 1;
	`, chatCategoryTable, categoryTable)

	if err := r.db.Select(&items, query, chatId); err != nil {
		return nil, err
	}

	return items, nil
}
