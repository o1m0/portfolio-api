package models

import "gorm.io/gorm"

type WorkCategory struct {
	gorm.Model
	WorkID     uint
	CategoryID uint
}
