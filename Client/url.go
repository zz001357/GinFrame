/**
* @Author: ZhangZe
* @Date: 2022/6/18
 */

package main

import (
	"GinFrame/Client/api"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/host"
	"google.golang.org/grpc"
)

var BlogServerAddr string
var ResumeServerAddr string
var PortfoliosServerAddr string
var RecordServerAddr string

func conn(ip string) *grpc.ClientConn {
	//创建一个grpc连接

	conn, errGrpc := grpc.Dial(ip, grpc.WithInsecure())
	if errGrpc != nil {
		panic(errGrpc)
	}
	return conn
}

func Server(r *gin.Engine) (*grpc.ClientConn, *grpc.ClientConn, *grpc.ClientConn, *grpc.ClientConn) {
	nInfo, _ := host.Info()
	if nInfo.OS == "windows" {
		BlogServerAddr = "127.0.0.1:8005"
		ResumeServerAddr = "127.0.0.1:8006"
		PortfoliosServerAddr = "127.0.0.1:8007"
		RecordServerAddr = "127.0.0.1:8008"
	} else {
		BlogServerAddr = "gve_blog:8005"
		ResumeServerAddr = "gve_resume:8006"
		PortfoliosServerAddr = "gve_portfolios:8007"
		RecordServerAddr = "gve_record:8008"
	}
	conn1 := conn(BlogServerAddr)
	conn2 := conn(ResumeServerAddr)
	conn3 := conn(PortfoliosServerAddr)
	conn4 := conn(RecordServerAddr)

	api.BlogView(r, conn1)
	api.ResumeView(r, conn2)
	api.PortfoliosView(r, conn3)
	api.RecordView(r, conn4)
	return conn1, conn2, conn3, conn4
}
