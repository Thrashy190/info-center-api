package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	IDusers        uint   `gorm:"primaryKey; not null; uniqueIndex" json:"id"`
	Username       string `gorm:"not null; uniqueIndex" json:"username"`
	Password       string `gorm:"not null" json:"password"`
	Name           string `gorm:"not null" json:"name"`
	FatherLastName string `gorm:"not null" json:"father_last_name"`
	MotherLastName string `gorm:"not null" json:"mother_last_name"`
}
