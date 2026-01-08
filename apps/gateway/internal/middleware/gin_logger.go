package middleware

import (
	"ChatServer/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// GinLogger 接收 gin 框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(c.Request.Context(), "gin request",
			logger.Int("status", c.Writer.Status()),
			logger.String("method", c.Request.Method),
			logger.String("path", path),
			logger.String("query", query),
			logger.String("ip", c.ClientIP()),
			logger.String("user-agent", c.Request.UserAgent()),
			logger.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			logger.Duration("cost", cost),
		)
	}
}

// GinRecovery recover 项目可能出现的 panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// TODO: 完善 panic 的日志记录，包括堆栈信息
				logger.Error(c.Request.Context(), "gin panic",
					logger.Any("error", err),
				)
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
