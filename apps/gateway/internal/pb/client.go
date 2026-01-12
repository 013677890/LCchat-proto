package pb

import (
	userpb "ChatServer/apps/user/pb"
	"context"
	"fmt"
	"time"

	"ChatServer/apps/gateway/internal/middleware"
	"ChatServer/pkg/logger"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// userServiceClient 用户服务的gRPC客户端
	userServiceClient userpb.UserServiceClient
	// userServiceConn 用户服务的gRPC连接
	userServiceConn *grpc.ClientConn
	// userServiceBreaker 用户服务的熔断器
	userServiceBreaker *gobreaker.CircuitBreaker
)

// gRPC 服务配置，定义重试策略
// 通过 JSON 配置实现自动重试
const retryPolicy = `{
	"methodConfig": [{
		"name": [{"service": "user.UserService"}],
		"waitForReady": true,
		"timeout": "2s",
		"retryPolicy": {
			"maxAttempts": 5,
			"initialBackoff": "0.1s",
			"maxBackoff": "1s",
			"backoffMultiplier": 2,
			"retryableStatusCodes": ["UNAVAILABLE", "DEADLINE_EXCEEDED", "UNKNOWN"]
		}
	}]
}`

// InitUserServiceClient 初始化用户服务gRPC客户端
// addr: 用户服务地址，格式为 "host:port"，例如 "localhost:9090"
func InitUserServiceClient(addr string) error {
	// 1. 初始化熔断器配置
	userServiceBreaker = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "user-service",
		MaxRequests: 3,                // 半开状态下最多允许 3 个请求尝试
		Interval:    15 * time.Second,  // 清除计数的时间间隔
		Timeout:     45 * time.Second, // 熔断器开启后多久尝试进入半开状态
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// 失败率超过 50% 且连续失败次数超过 5 次时触发熔断
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 5 && failureRatio >= 0.5
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			logger.Info(context.Background(), "熔断器状态变化",
				logger.String("name", name),
				logger.String("from", from.String()),
				logger.String("to", to.String()),
			)
		},
	})

	// 2. 建立gRPC连接并注入拦截器
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(retryPolicy), // 应用重试策略
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(4*1024*1024), // 4MB接收大小
		),
		// 注入熔断拦截器
		grpc.WithChainUnaryInterceptor(
			middleware.CircuitBreakerInterceptor(userServiceBreaker),
		),
	)
	if err != nil {
		return err
	}

	userServiceConn = conn
	userServiceClient = userpb.NewUserServiceClient(conn)

	return nil
}

// CloseUserServiceClient 关闭用户服务gRPC客户端
func CloseUserServiceClient() error {
	if userServiceConn != nil {
		if err := userServiceConn.Close(); err != nil {
			return err
		}
	}

	return nil
}

// GetUserServiceClient 获取用户服务gRPC客户端
func GetUserServiceClient() userpb.UserServiceClient {
	return userServiceClient
}

// Login 登录方法
// ctx: 上下文
// req: 登录请求
// 返回: 登录响应和错误
// 注意: gRPC 会根据配置的重试策略自动重试失败的请求
func Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	client := GetUserServiceClient()
	if client == nil {
		return nil, fmt.Errorf("user service client not initialized")
	}

	// 记录 gRPC 调用开始时间，用于计算耗时
	start := time.Now()

	// gRPC 会自动应用重试策略，无需手动重试
	resp, err := client.Login(ctx, req)

	// 计算耗时并记录到 Prometheus 指标
	duration := time.Since(start).Seconds()
	middleware.RecordGRPCRequest("user.UserService", "Login", duration, err)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
