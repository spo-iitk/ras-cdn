package models

import (
	"time"

	"gorm.io/gorm"
)

type Uploads struct {
	gorm.Model
	UserID string `gorm:"not null" json:"user_id"`
	Name   string `gorm:"not null" json:"name"`
	URL    string `json:"url"`
}

type Zips struct {
	gorm.Model
	Files      string    `json:"files"`
	OutFile    string    `json:"outfile"`
	AccessedAt time.Time `json:"accessed_at"`
}
