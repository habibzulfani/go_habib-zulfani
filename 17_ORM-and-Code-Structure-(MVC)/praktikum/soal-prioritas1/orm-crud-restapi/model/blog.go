package model

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model

	UserID string `json:"user-id" form:"user-id"`

	Title string `json:"title" form:"title"`

	Content string `json:"content" form:"content"`
}
