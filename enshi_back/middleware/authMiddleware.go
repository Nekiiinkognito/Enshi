package middleware

import (
	"enshi/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// token := c.GetHeader("Authorization")

		tokenFromCookies := c.Request.CookiesNamed("auth_cookie")[0].Value
		cookieClimes, err := auth.ValidateToken(tokenFromCookies)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error auth": err.Error()})
			c.Abort()
			return
		}

		// claims, err := auth.ValidateToken(token)
		// if err != nil {
		// 	c.IndentedJSON(http.StatusUnauthorized, gin.H{"error auth": err.Error()})
		// 	c.Abort()
		// 	return
		// }

		// Claims -> data stored in token
		c.Set(ContextUserId, cookieClimes["id"])
		c.Set(ContextTokenData, cookieClimes)
		c.Next()

	}
}
