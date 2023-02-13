package models

import "gorm.io/gorm"

type Careers struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey; not null; uniqueIndex; autoIncrement" json:"id"`
	Name string `gorm:"not null;" json:"name"`
}
