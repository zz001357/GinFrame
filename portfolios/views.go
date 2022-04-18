/**
* @Author: ZhangZe
* @Date: 2022/4/17 15:01
 */

package portfolios

import (
	"github.com/gin-gonic/gin"
)

func Photos(c *gin.Engine) {
	c.POST("/api/portfolios/getPhotosCategory", getPhotosCategory) //获取照片类别
}
