package routes

import (
	"golang2/controllers"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine{
	router := gin.Default()
	
// ===============Users Group =======================================
	Users_group := router.Group("/")

	Users_group.GET("", controllers.ViewUsers)
	Users_group.POST("/created", controllers.CreateUsers)
	Users_group.GET("/update/:id", controllers.Updateusers)
	Users_group.GET("/delete/:id", controllers.DeleteUsers)
	// ==============================================================
	
// ===============Member Group=======================================
	Member_group := router.Group("/member")

	Member_group.GET("", controllers.ViewMembers)
	Member_group.POST("/created", controllers.InsertMembers	)
	Member_group.GET("/update/:id", controllers.UpdateMember)
	Member_group.GET("/delete/:id", controllers.DeleteMember)
	return router
}