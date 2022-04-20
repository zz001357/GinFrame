/**
* @Author: ZhangZe
* @Date: 2022/4/17 16:29
 */

package blogs

import "github.com/gin-gonic/gin"

func Blogs(c *gin.Engine) {
	c.POST("/api/blogs/getBlogsCategory", getBlogsCategory) //获取博文类别

	/*文章内容接口相关*/
	c.POST("/api/blogs/getArticles", getArticles)             //获取文章类别
	c.POST("/api/blogs/getArticleContent", getArticleContent) //获取文章类别

}
