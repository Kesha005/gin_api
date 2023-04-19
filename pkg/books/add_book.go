package books

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kesha005/go_package/pkg/common/models"

)

type AddBookRequestBody struct {
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	Description string `json:"description" form:"description"`
}

func (h handler) AddBook(ctx *gin.Context) {

	body := AddBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, gin.H{
		"body":body,
	})
	var book models.Book
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusCreated, &book)
}
