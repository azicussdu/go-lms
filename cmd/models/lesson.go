package models

import "time"

type Lesson struct {
	ID           uint   `gorm:"primaryKey"`
	Title        string `gorm:"size:128;not null"`
	VideoUrl     string `gorm:"size:128"`
	Content      string `gorm:"type:text"`
	NumOfOrder   uint8
	IsVisible    bool `gorm:"not null;default:true"`
	TopicId      uint
	LessonType   string `gorm:"size:16;not null;default:theory"`
	NumOfSeconds uint
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    time.Time
}
