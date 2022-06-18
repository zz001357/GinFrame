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

	r.POST("/handle/blog/server", func(c *gin.Context) {
		/**
		 * @Name 记录服务请求
		 * @Param
		 * @Return
		 * @Date 2022/6/18 15:05
		 * @Author ZhangZe
		 **/
		ip, addr := common.IpUntil()                                                                                     //记录ip 和 地址
		reply, err := client.BlogServer(context.Background(), &pb.BlogRequest{ServerName: "blog", Ip: ip, IpAddr: addr}) //记录请求
		if err != nil {
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
		}
	})

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
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
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
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
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
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
		}
	})
}
