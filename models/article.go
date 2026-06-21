package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserID     uint
	Title      string
	Body       string
	Categories []Category `gorm:"many2many:article_categories;"`
}
