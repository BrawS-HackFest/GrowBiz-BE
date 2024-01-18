package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(50);not null"`
	Pict        string `json:"pict" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	Comments    []ArticleComment
}

type ArticleComment struct {
	gorm.Model
	Comment   string `json:"comment" gorm:"type:varchar(255);not null"`
	ArticleID uint
	UserID    string
}

type CommentPost struct {
	Comment string `json:"comment"`
}

type ArticleByIDResponse struct {
	ID          uint      `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Pict        string    `json:"pict"`
	Description string    `json:"description"`
	Comments    []Comment
}

type Comment struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	Comment   string    `json:"comment"`
	ArticleID uint      `json:"article_id"`
	User      UserReview
}
