package models

import "time"

type Topic struct {
	ID           uint   `gorm:"primaryKey"`
	Title        string `gorm:"size:128;not null"`
	Description  string `gorm:"size:256"`
	NumOfOrder   uint8
	IsVisible    bool `gorm:"not null;default:true"`
	NumOfSeconds uint
	NumOfLessons uint8
	CourseId     uint
	Lessons      []Lesson  `gorm:"foreignKey:TopicId"`
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    time.Time
}
