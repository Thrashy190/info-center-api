package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey; not null; uniqueIndex; autoIncrement" json:"id"`
	Title         string `gorm:"not null" json:"title"`
	Year          uint   `gorm:"not null" json:"year"`
	Clasification string `gorm:"not null" json:"clasification"`
	IsbnIssn      string `gorm:"not null" json:"isbn_issn"`
	Status        string `gorm:"not null" json:"status"`
	Typebook      string `gorm:"not null" json:"type"`
}
