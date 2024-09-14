package main

import (
	"context"
	"log"
	"time"

	_ "example.com/ningneng/init"
	"example.com/ningneng/internal/pb"
	"example.com/ningneng/pkg/global"
	"example.com/ningneng/pkg/interceptor/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	credTls, _ := credentials.NewClientTLSFromFile(
		global.RootPath+"/tools/key/test.pem",
		"*.ningneng.test",
	)

	// 设置签名私钥
	jwtAuth := auth.JwtAuthentication{[]byte("154a8b3aa89d3d4c49826f6dbbbe5542b5a9fbbb")}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(credTls))
	opts = append(opts, grpc.WithPerRPCCredentials(jwtAuth))

	// 连接server端，使用ssl加密通信
	conn, err := grpc.NewClient("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 调用登录
	resp, err := client.Login(ctx, &pb.LoginRequest{
		Username: "root",
		Password: "root",
	})

	if err != nil {
		if s, ok := status.FromError(err); ok {
			log.Fatalf("rpc error: code = %s, message = %s\n", s.Code().String(), s.Message())
		}
		log.Fatalf("login-response err: %v\n", err)
	}

	switch res := resp.Response.(type) {
	case *pb.LoginResponse_Account:
		log.Printf("登录成功：%s", res.Account)
	case *pb.LoginResponse_Error:
		log.Printf("登录失败：%s", res.Error)
	}

	// 调用getInfo
	resp01, err := client.GetInfo(ctx, &emptypb.Empty{})

	log.Printf("getInfo: %v", resp01)
}
