/**
* @Author: ZhangZe
* @Date: 2022/4/17 16:29
 */

package blogs

import (
	"GinFrame/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBlogsCategory(c *gin.Context) {
	/**
	 * @Name 获取博文类别
	 * @Param
	 * @Return
	 * @Date 2022/4/17 16:48
	 * @Author ZhangZe
	 **/
	sql := "select * from t_blogs_category where delete_time is null and is_show = '1'"
	data, err := common.ReadSql(sql)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}

func getArticleContent(c *gin.Context) {
	/**
	 * @Name 根据类别获取文章内容
	 * @Param blog_category_id
	 * @Return
	 * @Date 2022/4/18 14:22
	 * @Author ZhangZe
	 **/
	blogCategoryId := common.Params(c, "blog_category_id")
	sql := fmt.Sprintf("select * from t_blog_articles where is_show='1' and delete_time is null and blog_category_id='%s'", blogCategoryId)
	fmt.Println(sql)
	data, err := common.ReadSql(sql)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{Code: 1, Message: "查询失败", Data: err})
	} else {
		c.JSON(http.StatusOK, common.Response{Code: 0, Message: "查询成功", Data: data})
	}
}