/**
* @Author: ZhangZe
* @Date: 2022/4/10 16:09
 */

package router

import (
	"GinFrame/blogs"
	"GinFrame/portfolios"
	"GinFrame/resume"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Router(addr string) {
	r := gin.Default()
	r.Use(Cors())                           //开启中间件 允许使用跨域请求
	r.StaticFS("/imgs", http.Dir("./imgs")) //图片改为静态文件
	routers(r)
	if err := r.Run(addr); err != nil {
		log.Fatal("程序启动失败:", err)
	}

}
func routers(router *gin.Engine) {
	resume.Resume(router)
	portfolios.Photos(router)
	blogs.Blogs(router)
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
