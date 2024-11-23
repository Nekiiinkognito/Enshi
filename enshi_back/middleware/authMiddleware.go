package middleware

import (
	rest_api_stuff "enshi/REST_API_stuff"
	"enshi/auth"
	"enshi/global"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookies := c.Request.CookiesNamed("auth_cookie")
		if len(cookies) == 0 {
			rest_api_stuff.UnauthorizedAnswer(c, fmt.Errorf("no token provided"))
			c.Abort()
			return
		}

		tokenFromCookies := cookies[0].Value
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
