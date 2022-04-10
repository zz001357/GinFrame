/**
* @Author: ZhangZe
* @Date: 2022/4/10 16:04
 */

package admin

import (
	"github.com/gin-gonic/gin"
)

type ReqTest struct {
	UserName string `json:"user_name" binding:"required"` // 带校验方式
	Password string `json:"password"`
}

func Admin(c *gin.Engine) {
	c.GET("/login", login)
}
