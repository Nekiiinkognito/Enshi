package middleware

import (
	postspolicies "enshi/ABAC/postsPolicies"
	rest_api_stuff "enshi/REST_API_stuff"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		switch c.Request.Method {
		case "DELETE":
			c.Set("target", postspolicies.DELETE_POST)
		case "PUT":
			c.Set("target", postspolicies.UPDATE_POST)
		case "POST":
			c.Set("target", postspolicies.CREATE_POST)
		case "GET":
			c.Set("target", postspolicies.GET_POST)
		}

		isAllowed, err := postspolicies.PostsPolicies(c)

		if err != nil {
			rest_api_stuff.InternalErrorAnswer(c, err)
			c.Abort()
			return
		}

		if !isAllowed {
			rest_api_stuff.UnauthorizedAnswer(c, fmt.Errorf("you have no permission"))
			c.Abort()
			return
		}

		c.Next()
	}
}
