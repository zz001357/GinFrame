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
	"os"
	"path/filepath"
	"strings"
	"time"
)

var serviceName = "Life"

func getPhotosCategory(c *gin.Context) {
	/**
	 * @Name 获取照片类别
	 * @Param
	 * @Return
	 * @Date 2022/4/15 15:00
	 * @Author ZhangZe
	 **/
	sql := "SELECT a.id AS id,a.v_category_name AS v_category_name,b.v_photo_url FROM t_photos_category a LEFT JOIN " +
		"(select photo_category_id,v_photo_url from t_photos where is_first = 1 ) b ON a.id = b.photo_category_id  WHERE a.delete_time IS NULL " +
		"AND a.is_show = '1' ORDER BY created_time DESC"
	data, err := common.ReadSql(sql, common.Connection().GveLife)
	if err != nil {
		common.Failure(c, serviceName, err)
	} else {
		common.Success(c, serviceName, data)
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
	data, err := common.ReadSql(sql, common.Connection().GveLife)
	if err != nil {
		common.Failure(c, serviceName, err)
	} else {
		common.Success(c, serviceName, data)
	}
}

func uploadImg(c *gin.Context) {
	/**
	 * @Name 上传作品图片
	 * @Param	img_file	file	图片
	 * @Return
	 * @Date 2022/4/21 21:00
	 * @Author ZhangZe
	 **/
	img, err := c.FormFile("img_file")
	if err != nil {
		log.Println(err)
		common.Failure(c, serviceName, err)
	} else {
		fileExt := strings.ToLower(filepath.Ext(img.Filename)) //文件后缀名
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			common.Failure(c, serviceName, err)
		}
		folderPath := "img/portfolios/" + time.Now().Format("2006-01-02")
		_, errPath := os.Stat(folderPath) //Stat返回描述文件f的FileInfo类型值。如果出错，错误底层类型是*PathError。代表路径错误，可能是路径不存在
		if os.IsNotExist(errPath) {
			_ = os.MkdirAll(folderPath, os.ModePerm)
		}
		savePath := folderPath + "/" + time.Now().Format("20060102150405") + "_" + img.Filename
		uploadedErr := c.SaveUploadedFile(img, savePath)
		if uploadedErr == nil {
			common.Success(c, serviceName, savePath)
		} else {
			common.Failure(c, serviceName, err)
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
	err := common.InsertSql(sql, common.Connection().GveLife)
	if err != nil {
		common.Failure(c, serviceName, err)
	} else {
		common.Success(c, serviceName, nil)
	}
}
