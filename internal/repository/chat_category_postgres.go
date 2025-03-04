package repository

import (
	"fmt"
	"github.com/Koderbek/pocket_news_bot/internal/model"
	"github.com/jmoiron/sqlx"
)

type ChatCategoryPostgres struct {
	db *sqlx.DB
}

func NewChatCategoryPostgres(db *sqlx.DB) *ChatCategoryPostgres {
	return &ChatCategoryPostgres{db: db}
}

func (r *ChatCategoryPostgres) Add(chatId int64, categoryId int8, name string) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (chat_id, category_id, name)
		VALUES ($1, $2, $3)
		ON CONFLICT (chat_id, category_id)
			DO UPDATE SET active      = 'Y',
						  last_update = NOW();
	`, chatCategoryTable)

	_, err := r.db.Exec(query, chatId, categoryId, name)

	return err
}

func (r *ChatCategoryPostgres) Deactivate(chatId int64, categoryId int8) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET active      = 'N',
			last_update = NOW()
		WHERE chat_id = $1
		  AND category_id = $2;
	`, chatCategoryTable)

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
		WHERE cc.chat_id = $1 AND cc.active = 'Y';
	`, chatCategoryTable, categoryTable)

	if err := r.db.Select(&items, query, chatId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ChatCategoryPostgres) GetByCategoryId(categoryId int8) ([]model.ChatCategory, error) {
	var items []model.ChatCategory
	query := fmt.Sprintf(`
		SELECT cc.chat_id,
			   cc.name as chat_name,
			   cc.category_id,
			   c.code  as category_code,
			   c.name  as category_name
		FROM %s cc
				 INNER JOIN %s c on c.id = cc.category_id
		WHERE cc.category_id = $1 AND cc.active = 'Y';
	`, chatCategoryTable, categoryTable)

	if err := r.db.Select(&items, query, categoryId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ChatCategoryPostgres) HasChatCategory(chatId int64, categoryId int8) bool {
	var chatCategory model.ChatCategory
	query := fmt.Sprintf(
		"SELECT category_id FROM %s WHERE chat_id=$1 AND category_id=$2 AND active='Y'",
		chatCategoryTable,
	)

	err := r.db.Get(&chatCategory, query, chatId, categoryId)

	return err == nil && chatCategory.CategoryId == categoryId
}
