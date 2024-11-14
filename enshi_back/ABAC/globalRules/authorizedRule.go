package globalrules

import (
	"enshi/auth"
	"enshi/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizedRule(c *gin.Context) (bool, error) {
	tokenFromCookies := c.Request.CookiesNamed("auth_cookie")[0].Value
	cookieClimes, err := auth.ValidateToken(tokenFromCookies)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error auth": err.Error()})
		c.Abort()
		return false, err
	} else {
		c.Set(global.ContextUserId, cookieClimes["id"])
		c.Set(global.ContextTokenData, cookieClimes)
	}

	return true, nil
}
