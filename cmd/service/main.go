package main

import (
	"log"
	"net"

	_ "example.com/ningneng/init"
	"example.com/ningneng/internal/pb"
	"example.com/ningneng/pkg/global"
	"example.com/ningneng/pkg/interceptor/auth"
	l "example.com/ningneng/pkg/interceptor/log"
	"example.com/ningneng/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var (
		err  error
		opts = make([]grpc.ServerOption, 0, 8)
	)

	credTls, err := credentials.NewServerTLSFromFile(
		global.RootPath+"/tools/key/test.pem",
		global.RootPath+"/tools/key/test.key",
	)

	if err != nil {
		log.Fatalf("tls 错误：%v", err)
	}

	opts = append(opts, grpc.Creds(credTls)) // TLS
	//opts = append(opts, grpc.UnaryInterceptor(auth.UnaryServerAuthInterceptor())) // Auth

	opts = append(opts, grpc.ChainUnaryInterceptor(
		l.UnaryServerLoggerInterceptor(),  // Logger
		auth.UnaryServerAuthInterceptor(), // Auth
	))

	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("监听失败：%v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOrderManagementServer(grpcServer, &server.Authentication{})

	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("启动服务失败：%v", err)
	}
}
