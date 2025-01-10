package middleware

import (
	"blog/common/enum"
	"blog/common/response"
	"blog/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(enum.JwtTokenName)
		if token == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		c.Set(enum.ContextUserId, claims.UserId)
		c.Next()
	}
}
