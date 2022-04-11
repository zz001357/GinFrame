/**
* @Author: ZhangZe
* @Date: 2022/4/10 21:11
 */

package api

import "github.com/gin-gonic/gin"

func Api(c *gin.Engine) {
	c.GET("/api/getUserInfo", getUserInfo)     //获取基本信息
	c.GET("/api/getResumeInfo", getResumeInfo) //获取简历信息
}
