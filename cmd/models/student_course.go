package models

import (
	"time"
)

type StudentCourse struct {
	StudentID   uint      `gorm:"primaryKey"`
	CourseID    uint      `gorm:"primaryKey"`
	StartDate   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	EndDate     time.Time
	IsCompleted bool `gorm:"not null;default:false;index"`
	//Certificate needed info (so if he changes need to log it)
	Username string `gorm:"size:64"`
	Lastname string `gorm:"size:64"`
}
