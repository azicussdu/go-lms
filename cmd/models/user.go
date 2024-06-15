package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"size:64;unique;not null"`
	PasswordHash string `gorm:"size:128;not null"`
	Username     string `gorm:"size:64"`
	Lastname     string `gorm:"size:64"`
	Role         string `gorm:"size:32;not null;default:'go-user'"`
	PhoneNumber  string `gorm:"size:16"`
	BannedTill   time.Time
	Courses      []*Course `gorm:"many2many:student_courses;"`
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    time.Time
}
