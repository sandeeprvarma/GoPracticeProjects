package models

import "gorm.io/gorm"

type Book struct {
	ID     int64   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func MigrateBooks(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
