package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qiuhaonan/float-backend/internal/utils"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "未提供认证token")
			c.Abort()
			return
		}

		// Bearer Token 格式验证
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.ErrorResponse(c, http.StatusUnauthorized, "token格式错误")
			c.Abort()
			return
		}

		// 验证 Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "token无效或已过期")
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
