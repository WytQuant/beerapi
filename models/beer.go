package models

import "gorm.io/gorm"

type Beer struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Category string `gorm:"not null"`
	Detail   string `gorm:"not null"`
	Image    string `gorm:"not null"`
}
