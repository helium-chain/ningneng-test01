package server

import (
	"context"

	"example.com/ningneng/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// 编译检查 *Authentication 是否实现了 pb.OrderManagementServer 接口
var _ pb.OrderManagementServer = (*Authentication)(nil)

type Authentication struct {
	pb.UnimplementedOrderManagementServer
}

// Login 登录
func (a *Authentication) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// 验证登录逻辑
	if req.Username == "root" && req.Password == "root" {
		return &pb.LoginResponse{
			Response: &pb.LoginResponse_Account{Account: req.Username},
		}, status.New(codes.OK, "").Err()
	}

	return &pb.LoginResponse{
		Response: &pb.LoginResponse_Error{Error: "Unauthorized"},
	}, status.Errorf(codes.Unauthenticated, "登录失败")
}

func (a *Authentication) GetInfo(ctx context.Context, req *emptypb.Empty) (*pb.InfoResponse, error) {
	return &pb.InfoResponse{
		Account: "root",
		Name:    "张三",
		Age:     15,
	}, nil
}
