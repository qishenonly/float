package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/utils"
	"github.com/qiuhaonan/float-backend/pkg/logger"
)

// Recovery 错误恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("Panic recovered: %v", err)
				utils.ErrorResponse(c, http.StatusInternalServerError, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
