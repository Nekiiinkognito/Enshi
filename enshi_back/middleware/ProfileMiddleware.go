package middleware

import (
	profilepolicies "enshi/ABAC/ProfilePolicies"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func ProfileMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case "PUT":
			c.Set("target", profilepolicies.UPDATE_PROFILE)
		}

		isAllowed, errors := profilepolicies.ProfilePolicies(c)

		if rules.ShouldAbortRequest(c, isAllowed, errors) {
			c.Abort()
			return
		}

		c.Next()
	}
}
