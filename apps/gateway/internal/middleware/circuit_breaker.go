package middleware

import (
	"context"
	"fmt"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CircuitBreakerInterceptor 创建一个 gRPC 客户端一元拦截器，用于实现熔断保护
// cb: 针对该服务的熔断器实例
func CircuitBreakerInterceptor(cb *gobreaker.CircuitBreaker) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 使用 gobreaker 执行请求
		_, err := cb.Execute(func() (interface{}, error) {
			// 执行实际的 RPC 调用
			err := invoker(ctx, method, req, reply, cc, opts...)
			return nil, err
		})

		if err != nil {
			// 如果熔断器处于开启状态，直接返回 Unavailable 错误
			if err == gobreaker.ErrOpenState {
				return status.Error(codes.Unavailable, fmt.Sprintf("circuit breaker [%s] is open", cb.Name()))
			}
			// 其他错误（如 RPC 调用本身的错误）会由 gobreaker 记录并统计失败率
			return err
		}

		return nil
	}
}
