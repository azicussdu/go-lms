package models

import "time"

type Author struct {
	ID          uint      `gorm:"primaryKey"`
	AuthorName  string    `gorm:"size:64;not null"`
	Position    string    `gorm:"size:128"`
	WebsiteUrl  string    `gorm:"size:128"`
	Description string    `gorm:"size:512"`
	Courses     []*Course `gorm:"many2many:course_authors;"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   time.Time
}
