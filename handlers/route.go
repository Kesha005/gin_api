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
	routes.PUT("/update/:id",h.Update)
	routes.GET("/:id",h.GetById)
	routes.DELETE("/:id",h.Delete)


	routes1 := route.Group("/user")
	routes1.GET("/all",h.AllUser)
	routes1.POST("/store",h.StoreUser)
	routes1.PUT("/update/:id",h.Updateuser)
	routes1.GET("/:id",h.USerbyId)
	routes1.DELETE("/:id",h.DelUser)
}