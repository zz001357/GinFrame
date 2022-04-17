/**
* @Author: ZhangZe
* @Date: 2022/4/10 16:09
 */

package router

import (
	"GinFrame/resume"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Router(addr string) {
	r := gin.Default()
	r.Use(Cors()) //开启中间件 允许使用跨域请求
	routers(r)
	if err := r.Run(addr); err != nil {
		log.Fatal("程序启动失败:", err)

	}
}
func routers(router *gin.Engine) {
	resume.Api(router)
}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
