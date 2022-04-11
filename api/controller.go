/**
* @Author: ZhangZe
* @Date: 2022/4/10 21:11
 */

package api

import (
	"GinFrame/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getUserInfo(c *gin.Context) {
	/**
	 * @Name 获取基本信息
	 * @Param
	 * @Return
	 * @Date 2022/4/11 15:00
	 * @Author ZhangZe
	 **/
	sql := "select * from t_user_info"
	data, err := common.ReadSql(sql)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}

func getResumeInfo(c *gin.Context) {
	/**
	 * @Name 获取简历信息
	 * @Param
	 * @Return
	 * @Date 2022/4/11 15:01
	 * @Author ZhangZe
	 **/
	sql := "select * from t_user_info"
	data, err := common.ReadSql(sql)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}
