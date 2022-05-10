/**
* @Author: ZhangZe
* @Date: 2022/5/9 22:02
 */

package main

import (
	"GinFrame/ResumeServer/api"
	"GinFrame/ResumeServer/common"
	pb "GinFrame/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"runtime"
)

func main() {
	var Url string
	sysType := runtime.GOOS
	if sysType == "windows" {
		Url = common.Config().GetString("Url.devUrl")
	} else {
		Url = common.Config().GetString("Url.proUrl")
	}
	lis, err := net.Listen("tcp", Url+":8006")
	if err != nil {
		log.Fatalf("服务端连接出错，服务停止，错误：+%v", err)
	}

	//定义一个rpc的server
	server := grpc.NewServer()
	//注册服务
	pb.RegisterResumeServer(server, &api.Server{})
	//进行映射绑定
	reflection.Register(server)

	//启动服务
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("服务端连接出错，服务停止，错误：+%v", err)
	}
}
