package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString = tokenString[7:] // Remove "Bearer " prefix

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secret := os.Getenv("TOKEN_SECRET")
			return []byte(secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Ensure token validity
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		authorities := claims["authorities"].([]interface{})
		var roles []string
		for _, authority := range authorities {
			name, ok := authority.(map[string]interface{})["name"].(string)
			if ok {
				roles = append(roles, name)
			}
		}

		c.Set("roles", roles)
		// Token is valid, proceed with request
		c.Next()
	}
}
