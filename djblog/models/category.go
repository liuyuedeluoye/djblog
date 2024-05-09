package models

type Category struct {
	CategoryID   int64  `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_id"`
}
