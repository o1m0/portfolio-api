package models

import "gorm.io/gorm"

type ArticleCategory struct {
	gorm.Model
	ArticleID  uint
	CategoryID uint
}
