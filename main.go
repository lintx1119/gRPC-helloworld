package main

import (
	"context"
	"fmt"
	pb "github.com/lintx1119/gRPC-helloworld/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

//func (UnimplementedSayHelloServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
//}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("SayHello called")
	return &pb.HelloResponse{
		ResponseMsg: "hello" + req.RequestName,
	}, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("failed to serve:", err)
		return
	}
}
