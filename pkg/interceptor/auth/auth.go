package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 不需要认证的放行
var guards = map[string]struct{}{
	"/OrderManagement/login": {},
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid credentials")
)

func UnaryServerAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// 白名单跳过
		if _, ok := guards[info.FullMethod]; ok {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errMissingMetadata
		}

		authorization := md["authorization"]
		if len(authorization) < 1 {
			return nil, errInvalidToken
		}

		tokenString := strings.TrimPrefix(md["authorization"][0], "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("签名方法错误: %v", token.Header["alg"])
			}
			// 签名密钥
			return []byte("154a8b3aa89d3d4c49826f6dbbbe5542b5a9fbbb"), nil
		})

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok || !token.Valid {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		fmt.Println("token认证成功：", claims.ID)

		return handler(ctx, req)
	}
}
