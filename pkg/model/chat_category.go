package model

type ChatCategory struct {
	ChatId       int    `db:"chat_id"`
	ChatName     string `db:"chat_name"`
	CategoryId   int    `db:"category_id"`
	CategoryCode string `db:"category_code"`
	CategoryName string `db:"category_name"`
}
