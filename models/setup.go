package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=26221030 dbname=test port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database!")
	}
	err = database.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to migrate")
	}
	DB = database
}
