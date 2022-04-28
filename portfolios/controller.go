/**
* @Author: ZhangZe
* @Date: 2022/4/17 15:02
 */

package portfolios

import (
	"GinFrame/common"
	"fmt"
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
	data, err := common.ReadSql(sql, common.Connection().GoFrame)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}

func getPhotos(c *gin.Context) {
	/**
	 * @Name 根据类别获取图片
	 * @Param	photo_category_id	string	类别ID
	 * @Return
	 * @Date 2022/4/24 21:59
	 * @Author ZhangZe
	 **/
	photoCategoryID := common.Params(c, "photo_category_id")
	sql := fmt.Sprintf("select * from t_photos where photo_category_id ='%s'", photoCategoryID)
	data, err := common.ReadSql(sql, common.Connection().GoFrame)
	fmt.Println(data[0]["v_photo_url"])
	if err != nil {
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}

func uploadImg(c *gin.Context) {
	/**
	 * @Name 上传作品图片
	 * @Param	img_file	file	图片
	 * @Param	img_category	string	图片类别名
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

func saveImg(c *gin.Context) {
	/**
	 * @Name 保存图片
	 * @Param	v_photo_alt	string	图片描述
	 * @Param	v_photo_alt	string	图片链接
	 * @Param	photo_category_id	string	类别ID
	 * @Return
	 * @Date 2022/4/24 21:58
	 * @Author ZhangZe
	 **/
	photoAlt := common.Params(c, "v_photo_alt")
	photoUrl := common.Params(c, "v_photo_url")
	photoCategoryID := common.Params(c, "photo_category_id")
	uploadTime := time.Now().Format("2006-01-02 15:04:05") //上传时间

	sql := fmt.Sprintf("insert into t_photos(photo_category_id,v_photo_alt,v_photo_url,upload_time) values('%s','%s','%s','%s')", photoCategoryID, photoAlt, photoUrl, uploadTime)
	err := common.InsertSql(sql, common.Connection().GoFrame)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "同步失败！", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "同步成功！", Data: nil})
	}
}
