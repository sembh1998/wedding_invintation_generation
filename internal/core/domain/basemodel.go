package domain

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel is the base model for all models with gorm
type BaseModel struct {
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
