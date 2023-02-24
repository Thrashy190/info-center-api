package models

import (
	"gorm.io/gorm"
)

type Projects struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey; not null; uniqueIndex; autoIncrement" json:"id"`
	Year        uint    `json:"year"`
	NumControl  string  `gorm:"uniqueIndex" json:"num_control"`
	ProjectName string  `gorm:"not null; uniqueIndex" json:"project_name"`
	CareersID   uint    `json:"careers_id"`
	Careers     Careers `json:"careers"`
}
