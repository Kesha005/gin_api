package handlers

import (
	"gin_crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func (db datab) AllUser(ctx *gin.Context) {
	var users []models.UserModel

	result := db.Db.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(http.StatusOK,&users)
}

func (db datab) StoreUser(ctx *gin.Context) {
	body := UserRequest{}

	var user models.UserModel
	user.Name = body.Name
	user.Email = body.Email
	user.Address = body.Address

	db.Db.Create(&user)

	ctx.JSON(200, gin.H{
		"user": user,
	})

}

func (db datab) USerbyId(ctx *gin.Context) {
	var user models.UserModel
	id := ctx.Param("id")
	if res := db.Db.First(&user, id); res.Error != nil {
		ctx.JSON(404, gin.H{
			"error": res,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

func (db datab) Updateuser(ctx *gin.Context) {

}

func (db datab) DelUser(ctx *gin.Context) {
	var user models.UserModel
	id := ctx.Param("id")

	if res := db.Db.First(&user, id); res.Error != nil {
		ctx.JSON(403, gin.H{
			"err": res.Error,
		})
		return
	}
	db.Db.Delete(&user)
	ctx.JSON(200, gin.H{
		"msg": "succ",
	})
}
