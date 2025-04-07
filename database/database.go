package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBconn *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/newnew?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Database connection failed")
	}
	log.Println("Connection Sucessful")

	DBconn = db
}
