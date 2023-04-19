package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/Kesha005/book_crud/pkg/books"
	"github.com/Kesha005/book_crud/pkg/common/db"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	port := ":80"
	dbUrl := "root:@tcp(127.0.0.1:3306)/forgolang?charset=utf8mb4&parseTime=True&loc=Local"

	router := gin.Default()
	dbHandler := db.Init(dbUrl)
	books.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})
	router.Run(port)
}
