/**
* @Author: ZhangZe
* @Date: 2022/6/3 13:28
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

func RecordView(r *gin.Engine, conn *grpc.ClientConn) {
	//创建gRPC客户端
	client := pb.NewRecordClient(conn)

	r.POST("/handle/record/server", func(c *gin.Context) {
		/**
		 * @Name 记录服务请求
		 * @Param
		 * @Return
		 * @Date 2022/6/18 15:05
		 * @Author ZhangZe
		 **/
		ip, addr := common.IpUntil()                                                                                           //记录ip 和 地址
		reply, err := client.RecordServer(context.Background(), &pb.RecordRequest{ServerName: "record", Ip: ip, IpAddr: addr}) //记录请求
		if err != nil {
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
		}
	})

	r.POST("/handle/record/getOtherRecord", func(c *gin.Context) {
		/**
		 * @Name 获取记录
		 * @Param record_style 记录类型(用在哪里的)
		 * @Return
		 * @Date 2022/6/3 14:26
		 * @Author ZhangZe
		 **/
		// 调用方法
		param := common.Params(c, "record_style")
		reply, err := client.GetOtherRecord(context.Background(), &pb.OtherRecordRequest{Name: "获取记录", Message: "ok", Param: param})
		if err != nil {
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
		}
	})

	r.POST("/handle/record/getIterationRecord", func(c *gin.Context) {
		/**
		 * @Name 获取迭代记录
		 * @Param iteration_type 版本类型 (web、min-app)
		 * @Return
		 * @Date 2022/6/12 14:26
		 * @Author ZhangZe
		 **/
		// 调用方法
		param := common.Params(c, "iteration_type")
		reply, err := client.GetIterationRecord(context.Background(), &pb.IterationRecordRequest{Name: "获取记录", Message: "ok", Param: param})
		if err != nil {
			log.Println("Client端出错:", err)
			c.JSON(http.StatusOK, reply)
		} else {
			c.JSON(http.StatusOK, reply)
		}
	})
}
