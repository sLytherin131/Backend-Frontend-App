package database

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Sesuaikan username, password, dan nama database Anda
	dsn := "root:@tcp(127.0.0.1:3306)/panitia_webinar?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to the database")
}
