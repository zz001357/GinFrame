/**
* @Author: ZhangZe
* @Date: 2022/4/10 16:04
 */

package admin

import (
	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Engine) {
	c.GET("/admin/login", login)
}
