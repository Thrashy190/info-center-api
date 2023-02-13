package models

import "gorm.io/gorm"

type Tesis struct {
	gorm.Model
	ID uint `gorm:"primaryKey; not null; uniqueIndex; autoIncrement" json:"id"`
}
