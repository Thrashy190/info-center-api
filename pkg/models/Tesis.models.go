package models

import "gorm.io/gorm"

type Tesis struct {
	gorm.Model
	ID         uint    `gorm:"primaryKey; not null; uniqueIndex; autoIncrement" json:"id"`
	Title      string  `json:"title"`
	Year       uint    `json:"year"`
	Identifier uint    `json:"identifier"`
	CareersID  uint    `json:"careers_id"`
	Careers    Careers `json:"careers"`
}
