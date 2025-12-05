package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/pkg/logger"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		if raw != "" {
			path = path + "?" + raw
		}

		logger.Infof("[HTTP] %s %s %d %v",
			c.Request.Method,
			path,
			c.Writer.Status(),
			latency,
		)
	}
}
