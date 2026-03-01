package database

import (
	"exam-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, _ := gorm.Open(sqlite.Open("exam.db"), &gorm.Config{})
	db.AutoMigrate(&models.Question{})
	DB = db
}
