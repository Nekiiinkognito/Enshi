package middleware

import (
	postvotespolicies "enshi/ABAC/PostVotesPolicies"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func PostVotesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		switch c.Request.Method {
		case "DELETE":
			c.Set("target", postvotespolicies.DELETE_VOTE)

		case "POST":
			c.Set("target", postvotespolicies.CREATE_VOTE)

		case "GET":
			c.Set("target", postvotespolicies.READ_VOTE)
		}

		isAllowed, errors := postvotespolicies.PostVotePolicies(c)

		if rules.ShouldAbortRequest(c, isAllowed, errors) {
			c.Abort()
			return
		}

		c.Next()
	}
}
