package models

import (
	"time"
)

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:64;not null"`
	Description string    `gorm:"size:256"`
	Courses     []Course  `gorm:"foreignKey:CategoryId"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   time.Time
}
