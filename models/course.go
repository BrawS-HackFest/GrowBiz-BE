package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name            string `json:"name" gorm:"type:varchar(255);not null"`
	Desc            string `json:"desc" gorm:"type:text;not null"`
	Price           int    `json:"price" gorm:"type:int;not null"`
	Buyer           int    `json:"buyer" gorm:"type:int;not null"`
	Rating          int    `json:"rating" gorm:"type:int;not null"`
	Bab             string `json:"bab" gorm:"type:text;not null"`
	Link            string `json:"link" gorm:"type:varchar(255);not null"`
	Transactions    []Transaction
	Users           []User     `json:"users" gorm:"many2many:course_users;"`
	Categories      []Category `json:"categories" gorm:"many2many:course_categories;"`
	Reviews         []Review
	CourseMaterials []CourseMaterial
}

type CourseUser struct {
	CourseID uint   `json:"course_id" binding:"required"`
	UserID   string `json:"user_id" binding:"required" gorm:"type:varchar(255);not null"`
	IsRated  bool   `json:"is_rated" gorm:"type:bool;not null;default:false"`
}

type CourseUserPost struct {
	CourseID uint   `json:"course_id" binding:"required"`
	UserID   string `json:"user_id" binding:"required"`
}

type CourseMaterial struct {
	gorm.Model
	Title    string `json:"title" gorm:"type:varchar(255)"`
	Content  string `json:"content" gorm:"type:text;not null"`
	CourseID uint
}

type CourseResponseByID struct {
	ID     uint     `json:"id"`
	Name   string   `json:"name"`
	Desc   string   `json:"desc"`
	Price  int      `json:"price"`
	Buyer  int      `json:"buyer"`
	Rating int      `json:"rating"`
	Bab    []string `json:"bab"`
	Pict   string   `json:"pict"`
	Review []ReviewResult
}

type CourseResponse struct {
	ID     uint   `json:"id"`
	Pict   string `json:"pict"`
	Name   string `json:"name"`
	Buyer  int    `json:"buyer"`
	Price  int    `json:"price"`
	Rating int    `json:"rating"`
}
