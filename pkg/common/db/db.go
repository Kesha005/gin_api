package db

import (
	"log"
	"github.com/Kesha005/go_package/pkg/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB{
	db ,err := gorm.Open(mysql.Open(url),&gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Book{})
	return db
}