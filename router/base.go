package router

import (
	"GinFrame/controller"
	"github.com/gin-gonic/gin"
)

func base(router *gin.Engine) {
	baseUrl := router.Group("/api")
	{
		baseUrl.POST("/login", controller.Login)
	}
	return
}

