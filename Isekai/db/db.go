package db

import (
	"isekai/Isekai/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/isekai?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("could not connect MySQL")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
