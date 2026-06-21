package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	Articles []Article `gorm:"many2many:article_categories;"`
}
