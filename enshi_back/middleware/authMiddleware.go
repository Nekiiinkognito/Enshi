package middleware

import (
	"enshi/auth"
	"enshi/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenFromCookies := c.Request.CookiesNamed("auth_cookie")[0].Value
		cookieClimes, err := auth.ValidateToken(tokenFromCookies)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error auth": err.Error()})
			c.Abort()
			return
		}

		c.Set(global.ContextUserId, cookieClimes["id"])
		c.Set(global.ContextTokenData, cookieClimes)
		c.Next()

	}
}
