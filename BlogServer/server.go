/**
* @Author: ZhangZe
* @Date: 2022/5/4 23:41
 */

package main

import (
	"GinFrame/BlogServer/api"
	"GinFrame/BlogServer/common"
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
	lis, err := net.Listen("tcp", Url+":8005")
	if err != nil {
		log.Fatalf("服务端连接出错，服务停止，错误：+%v", err)
	}

	//定义一个rpc的server
	server := grpc.NewServer()
	//注册服务，相当与注册SayHello接口
	pb.RegisterBlogServer(server, &api.Server{})
	//进行映射绑定
	reflection.Register(server)

	//启动服务
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("服务端连接出错，服务停止，错误：+%v", err)
	}
}