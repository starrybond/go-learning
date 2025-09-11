package middleware

// JWT 认证中间件
// 任何需要登录才能访问的接口（发文章、删文章、发评论…）都会先经过它，验证通过后才能继续向后走，否则直接返回 401

import (
	"fmt"
	"net/http"

	"blog/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Println("Authorization:", c.GetHeader("Authorization"))
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			fmt.Println("parse fail:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
