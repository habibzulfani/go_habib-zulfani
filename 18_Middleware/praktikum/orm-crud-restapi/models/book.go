package models

import "github.com/jinzhu/gorm"

type Books struct {
	gorm.Model

	Title string `json:"title" form:"title"`

	Author string `json:"author" form:"author"`

	Publisher string `json:"publisher" form:"publisher"`
}
