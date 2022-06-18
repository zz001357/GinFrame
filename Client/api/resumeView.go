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
	//创建gRPC客户端
	client := pb.NewResumeClient(conn)

	r.POST("/handle/resume/server", func(c *gin.Context) {
		/**
		 * @Name 记录服务请求
		 * @Param
		 * @Return
		 * @Date 2022/6/18 15:05
		 * @Author ZhangZe
		 **/
		ip, addr := common.IpUntil()                                                                                            //记录ip 和 地址
		reply, err := client.RequestLogging(context.Background(), &pb.RequestParam{ServerName: "resume", Ip: ip, IpAddr: addr}) //记录请求
		if err != nil {
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
		}
	})
	//获取基本信息
	r.POST("/handle/resume/getBaseInfo", func(c *gin.Context) {
		/**
		 * @Name 获取基本信息
		 * @Param
		 * @Return
		 * @Date 2022/4/11 15:00
		 * @Author ZhangZe
		 **/
		reply, err := client.GetBaseInfo(context.Background(), &pb.BaseInfoRequest{Name: "获取基本信息", Message: "ok"}) // 调用方法
		if err != nil {
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
		}
	})

	//获取简历信息
	r.POST("/handle/resume/getResumeInfo", func(c *gin.Context) {
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
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
		}
	})
}
