package models

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	gorm.Model
	Comment  string `json:"comment" gorm:"type:text;not null"`
	Rating   uint   `json:"rating" gorm:"type:tinyint;not null"`
	UserId   string `json:"user_id" gorm:"type:varchar(255)"`
	CourseId uint   `json:"course_id" gorm:"type:int;not null"`
}

type ReviewPost struct {
	Comment string `json:"comment"`
	Rating  uint   `json:"rating"`
}

type ReviewResult struct {
	ID        uint       `json:"id"`
	UpdatedAt time.Time  `json:"updated_at"`
	Comment   string     `json:"comment"`
	Rating    uint       `json:"rating"`
	CourseID  uint       `json:"course_id"`
	User      UserReview `json:"user"`
}
