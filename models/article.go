package models

import "gorm.io/gorm"

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
