package models

import (
	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	ID                 uint          `gorm:"primaryKey; not null; uniqueIndex; autoIncrement" json:"id"`
	Description        string        `gorm:"not null" json:"description"`
	StateRequest       bool          `gorm:"default:false" json:"state_request"`
	RequestInformation []RequestInfo `json:"request_info"`
}
