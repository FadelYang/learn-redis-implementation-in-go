package middlewares

import (
	"errors"
	"net/http"
	"project-root/modules/auth/dto"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(accessSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errors": "missing authorization header",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errors": "missing authorization header",
			})
			return
		}

		tokenStr := parts[1]

		claims := &dto.AccessClaims{}

		token, err := jwt.ParseWithClaims(
			tokenStr,
			claims,
			func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing methid")
				}
				return accessSecret, nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errors": "invalid or expired token",
			})
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}
