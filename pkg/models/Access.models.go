package models

import "gorm.io/gorm"

type AccessStudents struct {
	gorm.Model
	Genre    string
	Semester uint
	CareerId uint
	Career   Careers
}

type AccessEmployees struct {
	gorm.Model
	Genre        string
	DepartmentId uint
	Department   Departments
}
