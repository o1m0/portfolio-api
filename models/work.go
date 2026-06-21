package models

import (
	"gorm.io/gorm"
)

type Work struct {
	gorm.Model
	Title       string
	Description string
	GithubURL   string
	DemoURL     string
	ImageURL    string
	Categories  []Category `gorm:"many2many:work_categories;"`
}
