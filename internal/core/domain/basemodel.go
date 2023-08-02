package domain

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel is the base model for all models with gorm
type BaseModel struct {
	ID        string         `json:"id" gorm:"primary_key;type:varchar(36);column:id"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
