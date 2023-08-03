package model

import "time"

type Article struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	PublishedAt time.Time `json:"publishedAt"`
}

type Articles struct {
	Articles []Article `json:"articles"`
}
