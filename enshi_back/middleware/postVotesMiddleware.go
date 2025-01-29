package middleware

import (
	postvotespolicies "enshi/ABAC/PostVotesPolicies"
	"enshi/ABAC/rules"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostVotesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := strings.Split(c.Request.URL.Path, "/")[1]
		switch c.Request.Method {
		case "DELETE":
			c.Set("target", postvotespolicies.DELETE_VOTE)

		case "POST":
			c.Set("target", postvotespolicies.CREATE_VOTE)

		case "GET":
			if a != "post-votes" {
				c.Set("target", postvotespolicies.READ_VOTE)
			} else {
				c.Set("target", "")
			}
		}

		isAllowed, errors := postvotespolicies.PostVotePolicies(c)

		if rules.ShouldAbortRequest(c, isAllowed, errors) {
			c.Abort()
			return
		}

		c.Next()
	}
}
