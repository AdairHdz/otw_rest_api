package entity

import (
	"time"
	"gorm.io/gorm"
)

type EntityUUID struct {
	ID        string `gorm:"primaryKey;size:36"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}