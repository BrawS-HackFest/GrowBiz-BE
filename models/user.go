package models

type User struct {
	Id              string     `json:"id" gorm:"primaryKey;unique;type:varchar(255);not null"`
	Username        string     `json:"username" gorm:"type:varchar(50)"`
	Email           string     `json:"email" gorm:"type:varchar(100);not null; unique"`
	Number          string     `json:"number" gorm:"type:varchar(15)"`
	Pict            string     `json:"pict" gorm:"type:varchar(255)"`
	Categories      []Category `json:"categories" gorm:"many2many:category_users;"`
	ArticleComments []ArticleComment
	Reviews         []Review
	Transactions    []Transaction
}

type UserCreate struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Number   string `json:"number"`
	Category []uint `json:"category"`
}

type UserReview struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
