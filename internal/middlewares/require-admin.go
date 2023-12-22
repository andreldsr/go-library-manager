package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		roles, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		}
		contains := utils.Contains(roles.([]string), "ROLE_ADMIN")
		if !contains {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		}
		c.Next()
	}
}
