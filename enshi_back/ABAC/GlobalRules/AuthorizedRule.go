package globalrules

import (
	"enshi/auth"
	"enshi/global"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizedRule(c *gin.Context) (bool, []error) {
	cookies := c.Request.CookiesNamed("auth_cookie")
	if len(cookies) == 0 {
		return false, []error{fmt.Errorf("no cookies provided")}
	}

	tokenFromCookies := cookies[0].Value
	cookieClimes, err := auth.ValidateToken(tokenFromCookies)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error auth": err.Error()})
		c.Abort()
		return false, []error{err}
	} else {
		c.Set(global.ContextUserId, cookieClimes["id"])
		c.Set(global.ContextTokenData, cookieClimes)
	}

	return true, nil
}
