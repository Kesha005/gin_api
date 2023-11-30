package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(route *gin.Engine, db *gorm.DB) {
	h := datab{
		Db: db,
	}
	routes := route.Group("/books")
	routes.GET("/all",h.All)
	routes.POST("/store",h.Store)
}