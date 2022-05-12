/**
* @Author: ZhangZe
* @Date: 2022/5/11 20:54
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

func PortfoliosView(r *gin.Engine, conn *grpc.ClientConn) {
	//创建RPC客户端
	client := pb.NewPortfoliosClient(conn)

	//获取照片类别
	r.POST("/api/portfolios/getImgCategory", func(c *gin.Context) {
		/**
		 * @Name 获取照片类别
		 * @Param
		 * @Return
		 * @Date 2022/4/15 15:00
		 * @Author ZhangZe
		 **/
		reply, err := client.GetImgCategory(context.Background(), &pb.ImgCategoryRequest{Name: "获取照片类别", Message: "ok"})
		if err != nil {
			log.Println("客户端出错:", err)
			c.JSON(http.StatusOK, common.Response{Code: 1, Message: "失败", Data: err})
		} else {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "成功", Data: reply})
		}
	})
	//根据类别Id获取图片
	r.POST("/api/portfolios/getImg", func(c *gin.Context) {
		/**
		 * @Name 根据类别Id获取图片
		 * @Param	img_category_id	string	类别ID
		 * @Return
		 * @Date 2022/4/24 21:59
		 * @Author ZhangZe
		 **/
		imgCategoryID := common.Params(c, "img_category_id")
		reply, err := client.GetImgByCategory(context.Background(), &pb.ImgRequest{Name: "获取照片类别", Message: "ok", ImgCategoryId: imgCategoryID})
		if err != nil {
			log.Println("客户端出错:", err)
			c.JSON(http.StatusOK, common.Response{Code: 1, Message: "失败", Data: err})
		} else {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "成功", Data: reply})
		}
	})
}