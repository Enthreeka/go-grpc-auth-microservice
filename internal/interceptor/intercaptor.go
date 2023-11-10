package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	panic("error")
}
