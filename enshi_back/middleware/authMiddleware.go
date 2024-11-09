package middleware

import (
	"enshi/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		// tokenFromCoolies := c.Request.CookiesNamed("auth_cookie")

		claims, err := auth.ValidateToken(token)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error auth": err.Error()})
			c.Abort()
			return
		}

		// Claims -> data stored in token
		c.Set("id", claims["id"])
		c.Set("claims", claims)
		c.Next()

	}
}
