package middleware

import (
	adminpolicies "enshi/ABAC/AdminPolicies"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAllowed, errors := adminpolicies.AdminPolicies(c)

		if rules.ShouldAbortRequest(c, isAllowed, errors) {
			c.Abort()
			return
		}

		c.Next()
	}
}
