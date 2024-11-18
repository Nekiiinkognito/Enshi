package middleware

import (
	postspolicies "enshi/ABAC/PostsPolicies"
	"enshi/ABAC/rules"
	"enshi/middleware/getters"

	"github.com/gin-gonic/gin"
)

func PostsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogId, _ := getters.GetInt64Param(c, "blog-id")
		postId, _ := getters.GetInt64Param(c, "post-id")

		switch c.Request.Method {
		case "DELETE":
			c.Set("target", postspolicies.DELETE_POST)
		case "PUT":
			if postId > 0 && blogId > 0 {
				c.Set("target", postspolicies.UPDATE_POST_BLOG)
			} else if postId > 0 {
				c.Set("target", postspolicies.UPDATE_POST)
			}

		case "POST":
			c.Set("target", postspolicies.CREATE_POST)
		case "GET":
			c.Set("target", postspolicies.GET_POST)
		}

		isAllowed, errors := postspolicies.PostsPolicies(c)

		if rules.ShouldAbortRequest(c, isAllowed, errors) {
			c.Abort()
			return
		}

		c.Next()
	}
}
