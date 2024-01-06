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
	Author string `json:"author"`
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
	var book models.Book

	book.Name = body.Name
	book.Author = body.Author

	if res := db.Db.Create(&book);res.Error!=nil{
		ctx.AbortWithError(http.StatusForbidden,res.Error)
		return 
	}
	ctx.JSON(http.StatusAccepted,&book)
	
}


func (db datab) Update(ctx *gin.Context){
	id := ctx.Param("id")
	body := BookRequest{}
	var book models.Book
	if err := ctx.BindJSON(&body); err!=nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return 
	}
	db.Db.First(&book,id)
	book.Name = body.Name
	book.Author = body.Author
}


func (db datab) GetById(ctx *gin.Context){
	id := ctx.Param("id")
	var book models.Book
	 db.Db.Find(&book,id)
	ctx.JSON(http.StatusOK, &book)
}

func (db datab) Delete(ctx *gin.Context){	
	id := ctx.Param("id")
	var book models.Book

	db.Db.First(&book,id)
	db.Db.Delete(&book)

	ctx.JSON(http.StatusOK,200)

}