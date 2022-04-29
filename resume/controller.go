/**
* @Author: ZhangZe
* @Date: 2022/4/10 21:11
 */

package resume

import (
	"GinFrame/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

var serviceName = "Resume"

func getUserInfo(c *gin.Context) {
	/**
	 * @Name 获取基本信息
	 * @Param
	 * @Return
	 * @Date 2022/4/11 15:00
	 * @Author ZhangZe
	 **/
	sql := "select * from t_user_info"
	data, err := common.ReadSql(sql, common.Connection().GveResume)
	if err != nil {
		common.Failure(c, serviceName, err)
	} else {
		common.Success(c, serviceName, data)
	}
}

func getResumeInfo(c *gin.Context) {
	/**
	 * @Name 获取简历信息
	 * @Param	search_key	string	模糊搜索关键字
	 * @Return
	 * @Date 2022/4/11 15:01
	 * @Author ZhangZe
	 **/
	param := common.Params(c, "search_key")
	var whereSql string
	if param != "" {
		whereSql = fmt.Sprint(" where  v_position like '%", param, "%' or v_word_content like '%", param, "%'")
	} else {
		whereSql = ""
	}
	sql := fmt.Sprint("select * from t_resume", whereSql)
	data, err := common.ReadSql(sql, common.Connection().GveResume)
	if err != nil {
		common.Failure(c, serviceName, err)
	} else {
		common.Success(c, serviceName, data)
	}
}
