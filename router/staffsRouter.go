package router

import (
	"GinFrame/controller"
	"github.com/gin-gonic/gin"
)

func staffsRouter(router *gin.Engine){
	staffsUrl := router.Group("/api")
	{
		staffsUrl.GET("/staffs", controller.GetStaffs)
	}
	return
}

