package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID             uint   `gorm:"primaryKey; not null; uniqueIndex; autoIncrement" json:"id"`
	Email          string `gorm:"not null; uniqueIndex" json:"email"`
	Password       string `gorm:"not null" json:"password"`
	Name           string `gorm:"not null" json:"name"`
	FatherLastName string `gorm:"not null" json:"father_last_name"`
	MotherLastName string `gorm:"not null" json:"mother_last_name"`
	EmployeeNumber string `gorm:"not null" json:"employee_number"`
}
