package models


import (
	"gorm.io/gorm"
)


type UserModel struct{
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
}