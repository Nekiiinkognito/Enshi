package middleware

import (
	bookmarkspolicies "enshi/ABAC/BookmarkPolicies"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func BookmarksMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		switch c.Request.Method {
		case "DELETE":
			c.Set("target", bookmarkspolicies.DELETE_BOOKMARK)

		case "POST":
			c.Set("target", bookmarkspolicies.CREATE_BOOKMARK)

		case "GET":
			c.Set("target", bookmarkspolicies.READ_BOOKMARK)
		}

		isAllowed, errors := bookmarkspolicies.BlogPolicies(c)

		if rules.ShouldAbortRequest(c, isAllowed, errors) {
			c.Abort()
			return
		}

		c.Next()
	}
}
