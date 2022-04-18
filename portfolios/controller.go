/**
* @Author: ZhangZe
* @Date: 2022/4/17 15:02
 */

package portfolios

import (
	"GinFrame/common"
	"github.com/gin-gonic/gin"
	"net/http"
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
	data, err := common.ReadSql(sql)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}
