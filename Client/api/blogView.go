/**
* @Author: ZhangZe
* @Date: 2022/5/5 20:36
 */

package api

import (
	"GinFrame/Client/common"
	pb "GinFrame/proto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func BlogView(r *gin.Engine, conn *grpc.ClientConn) {
	//创建RPC客户端
	client := pb.NewBlogClient(conn)
	//获取博文类别
	r.POST("/handle/blog/getBlogsCategory", func(c *gin.Context) {
		/**
		 * @Name 获取博文类别
		 * @Param
		 * @Return
		 * @Date 2022/5/7 16:48
		 * @Author ZhangZe
		 **/
		// 调用方法
		reply, err := client.GetBlogsCategory(context.Background(), &pb.BlogsCategoryRequest{Name: "获取博文类别", Message: "ok"})
		if err != nil {
			log.Println("客户端出错:", err)
			c.JSON(http.StatusOK, common.Response{Code: 1, Message: "失败", Data: err})
		} else {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "成功", Data: reply})
		}
	})

	/*---------------------------文章内容接口相关---------------------------*/

	//根据类别Id获取栏目信息
	r.POST("/handle/blog/getArticles", func(c *gin.Context) {
		/**
		 * @Name 根据类别获取文章内容
		 * @Param blog_category_id	string	博客类别id
		 * @Return
		 * @Date 2022/5/8 14:22
		 * @Author ZhangZe
		 **/
		blogCategoryId := common.Params(c, "blog_category_id")
		//调用方法
		reply, err := client.GetArticles(context.Background(), &pb.ArticlesRequest{Name: "根据类别Id获取栏目信息", Message: "ok", BlogCategoryId: blogCategoryId})
		if err != nil {
			log.Println("客户端出错:", err)
			c.JSON(http.StatusOK, common.Response{Code: 1, Message: "失败", Data: err})
		} else {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "成功", Data: reply})
		}
	})
	//根据id获取文章内容
	r.POST("/handle/blog/getArticleContent", func(c *gin.Context) {
		/**
		 * @Name 根据id获取文章内容
		 * @Param article_id	string	文章id
		 * @Return
		 * @Date 2022/4/18 14:22
		 * @Author ZhangZe
		 **/
		articleId := common.Params(c, "article_id")
		//调用方法
		reply, err := client.GetArticleContent(context.Background(), &pb.ArticlesContentRequest{Name: "根据类别Id获取栏目信息", Message: "ok", ArticleId: articleId})
		if err != nil {
			log.Println("客户端出错:", err)
			c.JSON(http.StatusOK, common.Response{Code: 1, Message: "失败", Data: err})
		} else {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "成功", Data: reply})
		}
	})
}
