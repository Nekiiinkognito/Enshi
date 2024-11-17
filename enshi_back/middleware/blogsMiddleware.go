package middleware

import (
	blogspolicies "enshi/ABAC/blogsPolicies"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func BlogsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		switch c.Request.Method {
		case "DELETE":
			c.Set("target", blogspolicies.DELETE_BLOG)
		case "PUT":
			c.Set("target", blogspolicies.UPDATE_BLOG)
		case "POST":
			c.Set("target", blogspolicies.CREATE_BLOG)
		case "GET":
			c.Set("target", blogspolicies.GET_BLOG)
		}

		isAllowed, errors := blogspolicies.BlogPolicies(c)

		if rules.ShouldAbortRequest(c, isAllowed, errors) {
			c.Abort()
			return
		}

		c.Next()
	}
}
