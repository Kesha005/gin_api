package handlers

import (
	"gin_crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type datab struct{
	Db *gorm.DB
}


type BookRequest struct{

	Name string `json:"name"`
	Author string `"author"`
}


func (db datab)All(ctx *gin.Context){
	var books []models.Book
	result := db.Db.Find(&books)
	if result.Error!=nil{
		panic(result.Error)
	}

	ctx.JSON(http.StatusOK,&books)

}

func (db datab) Store(ctx *gin.Context){
	body := BookRequest{}
	if  err := ctx.BindJSON(&body);err!=nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return 
	}
	var book models.Book

	book.Name = body.Name
	book.Author = body.Author

	if res := db.Db.Create(&book);res!=nil{
		ctx.AbortWithError(http.StatusForbidden,res.Error)
		return 
	}
	ctx.JSON(http.StatusAccepted,&book)
	
}


func (db datab) Update(ctx *gin.Context){

}


func (db datab) GetById(ctx *gin.Context){
	id := ctx.Param("id")
	var book models.Book
	result := db.Db.Find(&book,id)
	ctx.JSON(http.StatusOK, result)
}

func (db datab) Delete(ctx *gin.Context){

}