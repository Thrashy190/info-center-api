package models

import (
	"gorm.io/gorm"
)

type RequestInfo struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey; not null; uniqueIndex; autoIncrement" json:"id"`
	Year      uint   `json:"year"`
	Author    string `json:"author"`
	BookName  string `gorm:"not null" json:"book_name"`
	RequestID uint   `json:"request_id"`
}
