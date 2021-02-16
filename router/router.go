package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine)  {
	base(router)
	staffsRouter(router)
	return
}
