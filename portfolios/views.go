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
	c.POST("/api/portfolios/getPhotos", getPhotos)                 //根据类别或者图片id获取照片
	c.POST("/api/portfolios/uploadImg", UploadImg)                 //上传照片
}
