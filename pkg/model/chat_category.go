package model

type ChatCategory struct {
	ChatId       int64  `db:"chat_id"`
	ChatName     string `db:"chat_name"`
	CategoryId   int8   `db:"category_id"`
	CategoryCode string `db:"category_code"`
	CategoryName string `db:"category_name"`
}
