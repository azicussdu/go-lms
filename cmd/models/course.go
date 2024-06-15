package models

import "time"

type Course struct {
	ID                  uint   `gorm:"primaryKey"`
	Title               string `gorm:"size:128;not null"`
	Description         string `gorm:"type:text"`
	Slug                string `gorm:"size:64;not null;unique"`
	Price               uint   `gorm:"not null;default:0"`
	Rating              uint8
	Language            string `gorm:"size:16"`
	Level               string `gorm:"size:32;not null;default:beginner"`
	IsVisible           bool   `gorm:"not null;default:true"`
	EnrolledStudents    uint   `gorm:"not null;default:0"`
	NumOfSeconds        uint
	NumOfTheoryVideos   uint8
	NumOfPracticeVideos uint8
	NumOfMonths         uint8
	StartDate           time.Time
	EndDate             time.Time
	CategoryId          uint
	Topics              []Topic   `gorm:"foreignKey:CourseId"`
	Authors             []*Author `gorm:"many2many:course_authors;"`
	Users               []*User   `gorm:"many2many:student_courses;"`
	CreatedAt           time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt           time.Time
}
