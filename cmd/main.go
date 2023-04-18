package main

import (
	"github.com/Kesha005/go_package/common/db"
	"github.com/Kesha005/go_package/pkg/books"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main(){
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
	port :=viper.GET("PORT").(string)
	dbUrl := viper.GET("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dbUrl)
	books.RegisterRoutes(router ,dbHandler)

	router.GET("/",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"port":port,
			"dbUrl":dbUrl,
		})
	})
	router.Run(port)
}