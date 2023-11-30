package db

import (
	"gin_crud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

var DbHost string = "root:@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"

func Connect()*gorm.DB{
	
	DB, err := gorm.Open(mysql.Open(DbHost),&gorm.Config{})
	if err!=nil{
		panic(err)
	}
	DB.AutoMigrate(&models.Book{})
	

	return DB
}