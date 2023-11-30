package main

import (
	"gin_crud/db"
	"gin_crud/handlers"

	"github.com/gin-gonic/gin"
)


func main(){
	db_handler := db.Connect()
	router := gin.Default()
	handlers.RegisterRoutes(router,db_handler)

	router.GET("/", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "port":  80,
            
        })
    })

    router.Run(":80")
}