package middlewares

import (
	"net/http"
	"strings"

	"github.com/bhcoder23/gin-layout/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				`error`: "Authorization header can't be null",
			})
			return
		}
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != `bearer` {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				`error`: "Authorization header format error，must be: Bearer {token}",
			})
			return
		}
		token := parts[1]

		claims, err := utils.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				`error`:  "illegal token",
				`detail`: err.Error(),
			})
			return
		}
		c.Set(`user_id`, claims.UserID)
		c.Next()
	}
}
