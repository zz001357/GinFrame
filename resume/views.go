/**
* @Author: ZhangZe
* @Date: 2022/4/10 21:11
 */

package resume

import (
	"github.com/gin-gonic/gin"
)

func Resume(c *gin.Engine) {
	c.POST("/api/getUserInfo", getUserInfo)     //获取基本信息
	c.POST("/api/getResumeInfo", getResumeInfo) //获取简历信息
}
