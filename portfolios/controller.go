/**
* @Author: ZhangZe
* @Date: 2022/4/17 15:02
 */

package portfolios

import (
	"GinFrame/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getPhotosCategory(c *gin.Context) {
	/**
	 * @Name 获取照片类别
	 * @Param
	 * @Return
	 * @Date 2022/4/15 15:00
	 * @Author ZhangZe
	 **/
	sql := "select * from t_photos_category where delete_time is null and is_show = '1' order by created_time desc"
	data, err := common.ReadSql(sql, common.Connection("go_frame"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}

func getPhotos(c *gin.Context) {
	sql := "select * from t_photos"
	data, err := common.ReadSql(sql, common.Connection("go_frame"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}

func UploadImg(c *gin.Context) {
	/**
	 * @Name 上传作品图片
	 * @Param
	 * @Return
	 * @Date 2022/4/21 21:00
	 * @Author ZhangZe
	 **/
	img, err := c.FormFile("img_file")
	imgCategory := c.PostForm("img_category")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "上传失败", Data: err})
	} else {
		fileExt := strings.ToLower(filepath.Ext(img.Filename)) //文件后缀名
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "文件类型错误，上传失败！", Data: nil})
		}
		folderPath := filepath.Join("./imgs/", "portfolios", imgCategory)
		_, errPath := os.Stat(folderPath) //Stat返回描述文件f的FileInfo类型值。如果出错，错误底层类型是*PathError。代表路径错误，可能是路径不存在
		if os.IsNotExist(errPath) {
			_ = os.MkdirAll(folderPath, os.ModePerm)
		}
		savePath := folderPath + "\\" + time.Now().Format("20060102150405") + "_" + img.Filename
		uploadedErr := c.SaveUploadedFile(img, savePath)
		if uploadedErr == nil {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "上传成功！", Data: savePath})
		} else {
			log.Println(uploadedErr)
			c.JSON(http.StatusOK, common.Response{Code: 1, Message: "上传失败！", Data: uploadedErr})
		}
	}
}
