package database

import (
	"HackFest/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("init db failed,", err)
	}
	return DB
}

func MigrateDb(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Category{},
		&models.Review{},
		&models.CourseUser{},
		&models.CourseMaterial{},
		&models.Transaction{},
		&models.Article{},
		&models.ArticleComment{},
	)
	if err != nil {
		log.Fatal("Failed to migrate")
		return
	}
}
