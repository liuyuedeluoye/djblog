package models

import "time"

type Article struct {
	ArticleID int64 `db:"article_id" json:"article_id"`
	//CategoryID string    `db:"category_id" json:"category_id" binding:"required"`
	Title      string    `db:"title" json:"title" binding:"required"`
	Content    string    `binding:"required" db:"content" json:"content"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
