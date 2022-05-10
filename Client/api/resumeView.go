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

func ResumeView(r *gin.Engine, conn *grpc.ClientConn) {
	//创建RPC客户端
	client := pb.NewResumeClient(conn)
	//获取基本信息
	r.POST("/api/resume/getBaseInfo", func(c *gin.Context) {
		/**
		 * @Name 获取基本信息
		 * @Param
		 * @Return
		 * @Date 2022/4/11 15:00
		 * @Author ZhangZe
		 **/
		// 调用方法
		reply, err := client.GetBaseInfo(context.Background(), &pb.BaseInfoRequest{Name: "获取基本信息", Message: "ok"})
		if err != nil {
			log.Println("客户端出错:", err)
			c.JSON(http.StatusOK, common.Response{Code: 1, Message: "失败", Data: err})
		} else {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "成功", Data: reply})
		}
	})

	//获取简历信息
	r.POST("/api/resume/getResumeInfo", func(c *gin.Context) {
		/**
		 * @Name 获取简历信息
		 * @Param	search_key	string	模糊搜索关键字
		 * @Return
		 * @Date 2022/4/11 15:01
		 * @Author ZhangZe
		 **/
		param := common.Params(c, "search_key")
		//调用方法
		reply, err := client.GetResumeInfo(context.Background(), &pb.ResumeInfoRequest{Name: "获取简历信息", Message: "ok", SearchKey: param})
		if err != nil {
			log.Println("客户端出错:", err)
			c.JSON(http.StatusOK, common.Response{Code: 1, Message: "失败", Data: err})
		} else {
			c.JSON(http.StatusOK, common.Response{Code: 0, Message: "成功", Data: reply})
		}
	})
}
