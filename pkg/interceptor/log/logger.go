package log

import (
	"context"
	"time"

	"example.com/ningneng/pkg/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func UnaryServerLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// 记录请求日志信息
		global.Logger.Info("请求日志",
			zap.Time("TIME", time.Now()),
			zap.String("Method", info.FullMethod),
			zap.Any("Request", req),
		)

		response, err := handler(ctx, req)

		global.Logger.Info("响应日志",
			zap.Time("TIME", time.Now()),
			zap.String("Method", info.FullMethod),
			zap.Any("Response", response),
		)

		if err != nil {
			global.Logger.Error("错误日志",
				zap.Error(err),
			)
		}

		return response, err
	}
}
