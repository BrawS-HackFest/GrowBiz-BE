package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	OrderID       string `json:"order_id" gorm:"type:varchar(50);not null;unique"`
	TransactionID string `json:"transaction_id" gorm:"type:varchar(100);not null;"`
	Amount        int    `json:"amount" gorm:"type:int;not null"`
	Method        string `json:"method" gorm:"type:varchar(100);not null"`
	VaNumber      string `json:"va_number" gorm:"type:varchar(100);not null"`
	Status        string `json:"status" gorm:"type:varchar(100);not null"`
	CourseID      int
	UserID        string
}

type TransactionPost struct {
	Amount   int    `json:"amount" binding:"required"`
	Method   string `json:"method" binding:"required"`
	CourseID int    `json:"course_id" binding:"required"`
}

type TransactionByID struct {
	ID       uint   `json:"id"`
	Method   string `json:"method"`
	Amount   int    `json:"amount"`
	VANumber string `json:"va_number"`
	OrderID  string `json:"order_id"`
	Status   string `json:"status"`
}

type TransactionByUser struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Pict      string    `json:"pict"`
	Title     string    `json:"title"`
}
