package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			c.Abort()
			return
		}
		tokenString := parts[1]

		claims, err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token or expired token",
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()

	}
}

func RequiredRole(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "role is required",
			})
			c.Abort()
			return
		}
		roleString, ok := role.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "invalid role format",
			})
			c.Abort()
			return
		}
		authorized := false

		for _, requiredRole := range requiredRoles {
			if roleString == requiredRole {
				authorized = true
				break
			}

		}
		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "role is required",
			})
			c.Abort()
			return
		}
		c.Next()

	}
}
