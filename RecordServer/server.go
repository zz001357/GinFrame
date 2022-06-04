package main

import (
	"GinFrame/RecordServer/handle"
	pb "GinFrame/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("服务端连接出错，服务停止，错误：+%v", err)
	}
	//定义一个rpc的server
	server := grpc.NewServer()
	//注册服务
	pb.RegisterRecordServer(server, &handle.Server{})
	//进行映射绑定
	reflection.Register(server)
	//启动服务
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("服务端连接出错，服务停止，错误：+%v", err)
	}
}
