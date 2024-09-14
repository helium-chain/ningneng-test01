package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/credentials"
)

var _ credentials.PerRPCCredentials = JwtAuthentication{}

type JwtAuthentication struct {
	Key []byte
}

func (j JwtAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// 新创建一个令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        "example",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
	})

	// 使用secret签名并获得完整的编码令牌作为字符串
	tokenString, err := token.SignedString(j.Key)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"authorization": "Bearer " + tokenString,
	}, nil
}

func (JwtAuthentication) RequireTransportSecurity() bool {
	return true
}
