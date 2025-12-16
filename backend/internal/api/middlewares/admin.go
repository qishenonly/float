package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/utils"
)

// AdminMiddleware 管理员权限中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			utils.ErrorResponse(c, http.StatusForbidden, "没有权限")
			c.Abort()
			return
		}

		if roleStr, ok := role.(string); !ok || roleStr != "admin" {
			utils.ErrorResponse(c, http.StatusForbidden, "需要管理员权限")
			c.Abort()
			return
		}

		c.Next()
	}
}
